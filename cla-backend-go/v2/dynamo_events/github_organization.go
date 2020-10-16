// Copyright The Linux Foundation and each contributor to CommunityBridge.
// SPDX-License-Identifier: MIT

package dynamo_events

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/communitybridge/easycla/cla-backend-go/github"
	"github.com/communitybridge/easycla/cla-backend-go/github_organizations"
	log "github.com/communitybridge/easycla/cla-backend-go/logging"
	"github.com/communitybridge/easycla/cla-backend-go/utils"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

// GitHubOrgAddedEvent github repository added event
func (s *service) GitHubOrgAddedEvent(event events.DynamoDBEventRecord) error {
	f := logrus.Fields{
		"functionName": "GitHubOrgAddedEvent",
	}

	log.WithFields(f).Debug("processing event")
	var newGitHubOrg github_organizations.GithubOrganization
	err := unmarshalStreamImage(event.Change.NewImage, &newGitHubOrg)
	if err != nil {
		log.WithFields(f).Warnf("problem unmarshalling the new github organization model from the added event, error: %+v", err)
		return err
	}

	// If the branch protection value was updated from false to true....
	if newGitHubOrg.BranchProtectionEnabled {
		log.WithFields(f).Debug("branchProtectionEnabled - processing...")
		return s.enableBranchProtectionForGithubOrg(f, newGitHubOrg)
	}

	log.WithFields(f).Debug("no transition of branchProtectionEnabled - ignoring...")
	return nil
}

// GitHubOrgUpdatedEvent github repository updated event
func (s *service) GitHubOrgUpdatedEvent(event events.DynamoDBEventRecord) error {
	f := logrus.Fields{
		"functionName": "GitHubOrgUpdatedEvent",
	}

	log.WithFields(f).Debug("processing event")
	var newGitHubOrg, oldGitHubOrg github_organizations.GithubOrganization
	err := unmarshalStreamImage(event.Change.NewImage, &newGitHubOrg)
	if err != nil {
		log.WithFields(f).Warnf("problem unmarshalling the new github organization model from the updated event, error: %+v", err)
		return err
	}
	err = unmarshalStreamImage(event.Change.OldImage, &oldGitHubOrg)
	if err != nil {
		log.WithFields(f).Warnf("problem unmarshalling the old github organization model from the updated event, error: %+v", err)
		return err
	}

	// If the branch protection value was updated from false to true....
	if !oldGitHubOrg.BranchProtectionEnabled && newGitHubOrg.BranchProtectionEnabled {
		log.WithFields(f).Debug("transition of branchProtectionEnabled false => true - processing...")
		return s.enableBranchProtectionForGithubOrg(f, newGitHubOrg)
	}

	log.WithFields(f).Debug("no transition of branchProtectionEnabled false => true - ignoring...")
	return nil
}

// GitHubOrgDeletedEvent github repository deleted event
func (s *service) GitHubOrgDeletedEvent(event events.DynamoDBEventRecord) error {
	ctx := utils.NewContext()
	f := logrus.Fields{
		"functionName":   "GitHubOrgDeletedEvent",
		utils.XREQUESTID: ctx.Value(utils.XREQUESTID),
	}

	log.WithFields(f).Debug("processing event")
	var oldGitHubOrg github_organizations.GithubOrganization
	err := unmarshalStreamImage(event.Change.OldImage, &oldGitHubOrg)
	if err != nil {
		log.WithFields(f).WithError(err).Warnf("problem unmarshalling the old github organization model from the deleted event")
		return err
	}

	orgName := oldGitHubOrg.OrganizationName
	f["organizationName"] = orgName

	repoModels, err := s.repositoryService.GetRepositoriesByOrganizationName(ctx, orgName)
	if err != nil {
		log.WithFields(f).WithError(err).Warnf("problem loading repositories using org name: %s", orgName)
		return err
	}

	if len(repoModels) == 0 {
		log.WithFields(f).Debug("no repositories found for organization")
		return nil
	}

	log.WithFields(f).Debugf("disabling %d repositories for organization: %s", len(repoModels), orgName)
	for _, repo := range repoModels {
		disableErr := s.repositoryService.DisableRepository(ctx, repo.RepositoryID)
		if disableErr != nil {
			log.WithFields(f).WithError(disableErr).Warnf("problem disabling repository: %s", repo.RepositoryName)
		}
	}

	return nil
}

func (s *service) enableBranchProtectionForGithubOrg(f logrus.Fields, newGitHubOrg github_organizations.GithubOrganization) error {
	// Locate the repositories already saved under this organization
	log.WithFields(f).Debugf("loading repositories under the organization : %s", newGitHubOrg.OrganizationName)
	repos, err := s.repositoryService.GetRepositoriesByOrganizationName(context.Background(), newGitHubOrg.OrganizationName)
	if err != nil {
		log.WithFields(f).Warnf("problem locating repositories by organization name, error: %+v", err)
		return err
	}

	ctx := context.Background()
	log.WithFields(f).Debugf("creating a new GitHub client object for org: %s...", newGitHubOrg.OrganizationName)
	gitHubClient, clientErr := github.NewGithubAppClient(newGitHubOrg.OrganizationInstallationID)
	if clientErr != nil {
		return clientErr
	}

	branchProtectionRepo := github.NewBranchProtectionRepository(gitHubClient.Repositories, github.EnableBlockingLimiter())

	var eg errgroup.Group
	// a pool of 5 concurrent workers
	var workerTokens = make(chan struct{}, 5)
	for _, repo := range repos {
		// this is for goroutine local variables
		repo := repo
		// acquire a worker token to create a new goroutine
		workerTokens <- struct{}{}
		// Update the branch protection in a go routine...
		eg.Go(func() error {
			defer func() {
				<-workerTokens // release the workerToken
			}()
			log.WithFields(f).Debugf("enabling branch protection for repository: %s", repo.RepositoryName)

			log.WithFields(f).Debugf("looking up the default branch for the GitHub repository: %s...", repo.RepositoryName)
			defaultBranch, branchErr := branchProtectionRepo.GetDefaultBranchForRepo(ctx, newGitHubOrg.OrganizationName, repo.RepositoryName)
			if branchErr != nil {
				return branchErr
			}

			log.WithFields(f).Debugf("enabling branch protection on the default branch %s for the GitHub repository: %s...",
				defaultBranch, repo.RepositoryName)
			return branchProtectionRepo.EnableBranchProtection(ctx, newGitHubOrg.OrganizationName, repo.RepositoryName,
				defaultBranch, true, []string{utils.GitHubBotName}, []string{})
		})
	}

	// Wait for the go routines to finish
	log.WithFields(f).Debugf("waiting for %d repositories to complete...", len(repos))
	var branchProtectionErr error
	if loadErr := eg.Wait(); loadErr != nil {
		log.WithFields(f).Warnf("encountered branch protection setup error: %+v", loadErr)
		branchProtectionErr = loadErr
	}

	return branchProtectionErr
}

package main

import (
	"time"
)

// extracted from https://github.com/google/go-github
// just to parse data from github api

/*
Copyright (c) 2013 The go-github AUTHORS. All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are
met:

   * Redistributions of source code must retain the above copyright
notice, this list of conditions and the following disclaimer.
   * Redistributions in binary form must reproduce the above
copyright notice, this list of conditions and the following disclaimer
in the documentation and/or other materials provided with the
distribution.
   * Neither the name of Google Inc. nor the names of its
contributors may be used to endorse or promote products derived from
this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
"AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/

type RepositoryCommit struct {
	NodeID      *string   `json:"node_id,omitempty"`
	SHA         *string   `json:"sha,omitempty"`
	Commit      *Commit   `json:"commit,omitempty"`
	Author      *User     `json:"author,omitempty"`
	Committer   *User     `json:"committer,omitempty"`
	Parents     []*Commit `json:"parents,omitempty"`
	HTMLURL     *string   `json:"html_url,omitempty"`
	URL         *string   `json:"url,omitempty"`
	CommentsURL *string   `json:"comments_url,omitempty"`

	Stats *CommitStats  `json:"stats,omitempty"`
	Files []*CommitFile `json:"files,omitempty"`
}

type CommitStats struct {
	Additions *int `json:"additions,omitempty"`
	Deletions *int `json:"deletions,omitempty"`
	Total     *int `json:"total,omitempty"`
}

type CommitFile struct {
	SHA              *string `json:"sha,omitempty"`
	Filename         *string `json:"filename,omitempty"`
	Additions        *int    `json:"additions,omitempty"`
	Deletions        *int    `json:"deletions,omitempty"`
	Changes          *int    `json:"changes,omitempty"`
	Status           *string `json:"status,omitempty"`
	Patch            *string `json:"patch,omitempty"`
	BlobURL          *string `json:"blob_url,omitempty"`
	RawURL           *string `json:"raw_url,omitempty"`
	ContentsURL      *string `json:"contents_url,omitempty"`
	PreviousFilename *string `json:"previous_filename,omitempty"`
}

type Commit struct {
	SHA          *string                `json:"sha,omitempty"`
	Author       *CommitAuthor          `json:"author,omitempty"`
	Committer    *CommitAuthor          `json:"committer,omitempty"`
	Message      *string                `json:"message,omitempty"`
	Tree         *Tree                  `json:"tree,omitempty"`
	Parents      []*Commit              `json:"parents,omitempty"`
	Stats        *CommitStats           `json:"stats,omitempty"`
	HTMLURL      *string                `json:"html_url,omitempty"`
	URL          *string                `json:"url,omitempty"`
	Verification *SignatureVerification `json:"verification,omitempty"`
	NodeID       *string                `json:"node_id,omitempty"`

	CommentCount *int `json:"comment_count,omitempty"`
}

type CommitAuthor struct {
	Date  *time.Time `json:"date,omitempty"`
	Name  *string    `json:"name,omitempty"`
	Email *string    `json:"email,omitempty"`

	Login *string `json:"username,omitempty"`
}

type User struct {
	Login                   *string    `json:"login,omitempty"`
	ID                      *int64     `json:"id,omitempty"`
	NodeID                  *string    `json:"node_id,omitempty"`
	AvatarURL               *string    `json:"avatar_url,omitempty"`
	HTMLURL                 *string    `json:"html_url,omitempty"`
	GravatarID              *string    `json:"gravatar_id,omitempty"`
	Name                    *string    `json:"name,omitempty"`
	Company                 *string    `json:"company,omitempty"`
	Blog                    *string    `json:"blog,omitempty"`
	Location                *string    `json:"location,omitempty"`
	Email                   *string    `json:"email,omitempty"`
	Hireable                *bool      `json:"hireable,omitempty"`
	Bio                     *string    `json:"bio,omitempty"`
	TwitterUsername         *string    `json:"twitter_username,omitempty"`
	PublicRepos             *int       `json:"public_repos,omitempty"`
	PublicGists             *int       `json:"public_gists,omitempty"`
	Followers               *int       `json:"followers,omitempty"`
	Following               *int       `json:"following,omitempty"`
	CreatedAt               *Timestamp `json:"created_at,omitempty"`
	UpdatedAt               *Timestamp `json:"updated_at,omitempty"`
	SuspendedAt             *Timestamp `json:"suspended_at,omitempty"`
	Type                    *string    `json:"type,omitempty"`
	SiteAdmin               *bool      `json:"site_admin,omitempty"`
	TotalPrivateRepos       *int       `json:"total_private_repos,omitempty"`
	OwnedPrivateRepos       *int       `json:"owned_private_repos,omitempty"`
	PrivateGists            *int       `json:"private_gists,omitempty"`
	DiskUsage               *int       `json:"disk_usage,omitempty"`
	Collaborators           *int       `json:"collaborators,omitempty"`
	TwoFactorAuthentication *bool      `json:"two_factor_authentication,omitempty"`
	Plan                    *Plan      `json:"plan,omitempty"`
	LdapDn                  *string    `json:"ldap_dn,omitempty"`

	URL               *string `json:"url,omitempty"`
	EventsURL         *string `json:"events_url,omitempty"`
	FollowingURL      *string `json:"following_url,omitempty"`
	FollowersURL      *string `json:"followers_url,omitempty"`
	GistsURL          *string `json:"gists_url,omitempty"`
	OrganizationsURL  *string `json:"organizations_url,omitempty"`
	ReceivedEventsURL *string `json:"received_events_url,omitempty"`
	ReposURL          *string `json:"repos_url,omitempty"`
	StarredURL        *string `json:"starred_url,omitempty"`
	SubscriptionsURL  *string `json:"subscriptions_url,omitempty"`

	TextMatches []*TextMatch `json:"text_matches,omitempty"`

	Permissions map[string]bool `json:"permissions,omitempty"`
	RoleName    *string         `json:"role_name,omitempty"`
}

type Tree struct {
	SHA     *string      `json:"sha,omitempty"`
	Entries []*TreeEntry `json:"tree,omitempty"`

	Truncated *bool `json:"truncated,omitempty"`
}

type TreeEntry struct {
	SHA     *string `json:"sha,omitempty"`
	Path    *string `json:"path,omitempty"`
	Mode    *string `json:"mode,omitempty"`
	Type    *string `json:"type,omitempty"`
	Size    *int    `json:"size,omitempty"`
	Content *string `json:"content,omitempty"`
	URL     *string `json:"url,omitempty"`
}

type SignatureVerification struct {
	Verified  *bool   `json:"verified,omitempty"`
	Reason    *string `json:"reason,omitempty"`
	Signature *string `json:"signature,omitempty"`
	Payload   *string `json:"payload,omitempty"`
}

type Timestamp struct {
	time.Time
}

type Match struct {
	Text    *string `json:"text,omitempty"`
	Indices []int   `json:"indices,omitempty"`
}

type TextMatch struct {
	ObjectURL  *string  `json:"object_url,omitempty"`
	ObjectType *string  `json:"object_type,omitempty"`
	Property   *string  `json:"property,omitempty"`
	Fragment   *string  `json:"fragment,omitempty"`
	Matches    []*Match `json:"matches,omitempty"`
}

type Plan struct {
	Name          *string `json:"name,omitempty"`
	Space         *int    `json:"space,omitempty"`
	Collaborators *int    `json:"collaborators,omitempty"`
	PrivateRepos  *int    `json:"private_repos,omitempty"`
	FilledSeats   *int    `json:"filled_seats,omitempty"`
	Seats         *int    `json:"seats,omitempty"`
}

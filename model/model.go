package model

// GitInfo app git info
// {
// 	"version": "@pkg_version",
// 	"tag": "@pkg_branch_alias_version",
// 	"releaseAt": "@pkg_release_date"
// }
type GitInfo struct {
	Tag       string `json:"tag" description:"get tag name"`
	Version   string `json:"version" description:"git repo version."`
	ReleaseAt string `json:"releaseAt" description:"latest commit date"`
}

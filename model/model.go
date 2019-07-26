package model

// GitInfoData app git info
// {
// 	"version": "@pkg_version",
// 	"tag": "@pkg_branch_alias_version",
// 	"releaseAt": "@pkg_release_date"
// }
type GitInfoData struct {
	Tag       string `json:"tag" description:"get tag name"`
	Version   string `json:"version" description:"git repo version."`
	ReleaseAt string `json:"releaseAt" description:"latest commit date"`
}

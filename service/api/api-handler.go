package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	rt.router.POST("/session", rt.wrap(rt.doLogin))
	rt.router.GET("/users/:userid/profile/:username", rt.wrap(rt.getUserProfile))
	rt.router.GET("/users/:userid/profile/:username/", rt.wrap(rt.getUserProfileStream))
	rt.router.GET("/users", rt.wrap(rt.getMyStream))
	rt.router.PUT("/users/:userid/changeUsername", rt.wrap(rt.setMyUserName))
	rt.router.POST("/users/:userid/photos", rt.wrap(rt.uploadPhoto))
	rt.router.GET("/photos/:photoid", rt.wrap(rt.showPhoto))
	rt.router.DELETE("/photos/:photoid", rt.wrap(rt.deletePhoto))
	rt.router.GET("/photos/:photoid/comments/", rt.wrap(rt.listComments))
	rt.router.POST("/photos/:photoid/comments/", rt.wrap(rt.commentPhoto))
	rt.router.GET("/photos/:photoid/comments/:commentid", rt.wrap(rt.getComment))
	rt.router.DELETE("/photos/:photoid/comments/:commentid", rt.wrap(rt.uncommentPhoto))
	rt.router.GET("/photos/:photoid/likes/", rt.wrap(rt.listLikes))
	rt.router.GET("/photos/:photoid/likes/:userid", rt.wrap(rt.getLike))
	rt.router.POST("/photos/:photoid/likes/:userid", rt.wrap(rt.likePhoto))
	rt.router.DELETE("/photos/:photoid/likes/:userid", rt.wrap(rt.unlikePhoto))
	rt.router.GET("/users/:userid/following/", rt.wrap(rt.listFollowed))
	rt.router.GET("/users/:userid/following/:useridfollowed", rt.wrap(rt.getFollowed))
	rt.router.PUT("/users/:userid/following/:useridtofollow", rt.wrap(rt.followUser))
	rt.router.DELETE("/users/:userid/following/:useridtofollow", rt.wrap(rt.unfollowUser))
	rt.router.GET("/users/:userid/banned/", rt.wrap(rt.listBanned))
	rt.router.PUT("/users/:userid/banned/:useridtoban", rt.wrap(rt.banUser))
	rt.router.DELETE("/users/:userid/banned/:useridtoban", rt.wrap(rt.unbanUser))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}

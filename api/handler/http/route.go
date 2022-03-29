package http

func (sv *Server) registerRoutes() {
	public := sv.server.Group("/api/v1/public")
	private := sv.server.Group("/api/v1/private")
	internal := sv.server.Group("/api/v1/internal")
	public.GET("/ad-listing/:ad_id", sv.GetAdListing)
	private.GET("/ad-listing/:ad_id", sv.GetAdListing)
	internal.GET("/ad-listing/:ad_id", sv.GetAdListing)
}

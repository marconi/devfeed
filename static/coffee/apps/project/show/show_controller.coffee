define [
  "devfeed",
  "common_view",
  "project_show_view"
], (Devfeed, CommonView, ProjectShowView) ->

  Devfeed.module "ProjectApp.Show", (Show, Devfeed, Backbone, Marionette, $, _) ->

    Show.Controller =
      showProject: (id) ->
        # show preloader first while project is being loaded
        preloaderView = new CommonView.Preloader()
        Devfeed.contentRegion.show(preloaderView)

        fetchingProject = Devfeed.request("project:entity", id)
        $.when(fetchingProject).done (project) ->
          sidebarView = new ProjectShowView.Sidebar
          chatinfoView = new ProjectShowView.Chatinfo
          chatboxView = new ProjectShowView.Chatbox
          projectShowView = new ProjectShowView.Show
            model: project
          Devfeed.contentRegion.show(projectShowView)
          projectShowView.sidebarRegion.show(sidebarView)
          projectShowView.chatinfoRegion.show(chatinfoView)
          projectShowView.chatboxRegion.show(chatboxView)

          # render stories
          storiesView = new ProjectShowView.Stories
            collection: project.get("stories")
          sidebarView.storiesRegion.show(storiesView)

  return Devfeed.ProjectApp.Show.Controller

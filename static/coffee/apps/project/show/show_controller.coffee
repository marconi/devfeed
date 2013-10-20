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
            model: project
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
          storiesView.on "stories:more", ->
            fetchingStories = Devfeed.request("project:stories:more", project.get("id"))
            $.when(fetchingStories).done ->
              storiesView.triggerMethod("more:stories")
          storiesView.on "filters:changed", (filters) ->
            filteringStories = Devfeed.request("project:stories:filter", project.get("id"), filters)
            $.when(filteringStories).done ->
              # storiesView.triggerMethod("stories:filtered")
              sidebarView.triggerMethod("stories:filtered")
          sidebarView.storiesRegion.show(storiesView)

          # render find story form
          findStoryView = new ProjectShowView.FindStory
          findStoryView.on "settings:shown", ->
            sidebarView.triggerMethod("settings:shown")
          findStoryView.on "settings:hidden", ->
            sidebarView.triggerMethod("settings:hidden")
          findStoryView.on "filters:changed", (filters) ->
            sidebarView.triggerMethod("filters:changed", filters)
          sidebarView.findStoryRegion.show(findStoryView)

  return Devfeed.ProjectApp.Show.Controller

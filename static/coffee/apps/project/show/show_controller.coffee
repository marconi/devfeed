define [
  "devfeed",
  "common_view",
  "project_show_view"
], (Devfeed, CommonView, ProjectShowView) ->

  Devfeed.module "ProjectApp.Show", (Show, Devfeed, Backbone, Marionette, $, _) ->

    renderStories = (sidebarView, filters, project) ->
      storiesView = new ProjectShowView.Stories
        collection: project.get("stories")
      storiesView.on "stories:more", ->
        fetchingStories = Devfeed.request("project:stories", project.get("id"), filters, false, false)
        $.when(fetchingStories).done ->
          storiesView.triggerMethod("more:stories")
      sidebarView.storiesRegion.show(storiesView)
      return storiesView

    bindOnStoriesRendered = (storiesView, findStoryView) ->
      storiesView.on "stories:more:rendered", ->
        findStoryView.bindLiveFilter()

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
          storiesView = renderStories(sidebarView, [], project)

          # render find story form
          findStoryView = new ProjectShowView.FindStory
          bindOnStoriesRendered(storiesView, findStoryView)

          # handle triggers
          findStoryView.on "settings:shown", ->
            sidebarView.triggerMethod("settings:shown")
          findStoryView.on "settings:hidden", ->
            sidebarView.triggerMethod("settings:hidden")
          findStoryView.on "filters:changed", (filters) ->
            # show preloader first
            filterPreloaderView = new ProjectShowView.FilterPreloader
            sidebarView.storiesRegion.show(filterPreloaderView)

            # clear out existing stories
            project.get("stories").reset()

            # then fetch filtered stories
            fetchingStories = Devfeed.request("project:stories", project.get("id"), filters, true, true)
            $.when(fetchingStories).done ->
              storiesView = renderStories(sidebarView, findStoryView.filters, project)
              bindOnStoriesRendered(storiesView, findStoryView)
              if not $("#find-story .settings").hasClass("hide")
                storiesView.$el.addClass("settings-shown")
          sidebarView.findStoryRegion.show(findStoryView)

          # subscribe to project chat room
          Devfeed.execute("ws:project:subscribe", id)


  return Devfeed.ProjectApp.Show.Controller

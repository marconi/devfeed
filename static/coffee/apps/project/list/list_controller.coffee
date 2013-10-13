define [
  "devfeed",
  "common_view",
  "project",
  "project_list_view"
], (Devfeed, CommonView, Proj, ProjectListView) ->

  Devfeed.module "ProjectApp.List", (List, Devfeed, Backbone, Marionette, $, _) ->

    List.Controller =
      listProjects: ->
        # show preloader first while projects are being loaded
        preloaderView = new CommonView.Preloader()
        Devfeed.contentRegion.show(preloaderView)

        loadingProjects = Devfeed.request("project:entities")
        $.when(loadingProjects).done (projects) ->
          # console.log projects
          projectListView = new ProjectListView.List
            collection: projects
          Devfeed.contentRegion.show(projectListView)

  return Devfeed.ProjectApp.List.Controller

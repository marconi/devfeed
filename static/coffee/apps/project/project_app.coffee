define [
  "devfeed",
  "chat",
  "project_list_controller",
  "project_show_controller"
], (Devfeed, Chat, ProjectListController, ProjectShowController) ->

  Devfeed.module "ProjectApp", (ProjectApp, Devfeed, Backbone, Marionette, $, _) ->

    class ProjectApp.Router extends Marionette.AppRouter
      appRoutes:
        "projects": "listProjects"
        "projects/:id": "showProject"

    API =
      listProjects: ->
        Devfeed.isLoggedIn -> ProjectListController.listProjects()
      showProject: (id) ->
        Devfeed.isLoggedIn -> ProjectShowController.showProject(id)

    Devfeed.on "projects:list", ->
      Devfeed.navigate("projects")
      API.listProjects()

    Devfeed.on "project:show", (id) ->
      Devfeed.navigate("projects/" + id)
      API.showProject(id)

    Devfeed.addInitializer ->
      new ProjectApp.Router
        controller: API

  return Devfeed.ProjectApp

define ["devfeed"], (Devfeed) ->

  Devfeed.module "Entities", (Entities, Devfeed, Backbone, Marionette, $, _) ->

    Entities.Proj = {}

    class Entities.Proj.Project extends Backbone.Model
      defaults:
        id: null
        name: null
      urlRoot: "/api/projects"
      parse: (response, options) ->
        if response.s == 200
          return response.d

    class Entities.Proj.Projects extends Backbone.Collection
      model: Entities.Proj.Project
      url: "/api/projects"
      parse: (response, options) ->
        if response.s == 200
          return response.d

    projects = new Entities.Proj.Projects((Projects? and Projects) or null)

    API =
      getProject: (id) ->
        defer = $.Deferred()
        project = new Entities.Proj.Project id: id
        project.fetch
          success: (data) ->
            defer.resolve data
          error: (data) ->
            defer.resolve undefined
        return defer.promise()
      getProjects: ->
        defer = $.Deferred()
        if projects.length == 0
          projects.fetch
            success: (collection, response, options) ->
              if response.s == 200
                projects.reset response.d
              defer.resolve projects
            error: (collection, response, options) ->
              defer.resolve null
        else
          defer.resolve projects
        return defer.promise()

    Devfeed.reqres.setHandler "project:entity", (id) ->
      return API.getProject(id)

    Devfeed.reqres.setHandler "project:entities", ->
      return API.getProjects()

  return Devfeed.Entities.Proj

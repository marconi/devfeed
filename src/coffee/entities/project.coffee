define ["devfeed", "common_model"], (Devfeed, CommonModel) ->

  Devfeed.module "Entities", (Entities, Devfeed, Backbone, Marionette, $, _) ->

    Entities.Proj = {}

    class Entities.Proj.Task extends Backbone.Model
      defaults:
        oid: null
        id: null
        position: null
        description: null
        complete: false
        created_at: null
        updated_at: null

    class Entities.Proj.Tasks extends Backbone.Collection
      model: Entities.Proj.Task
      comparator: (task) ->
        return task.get("position")

    class Entities.Proj.Story extends Backbone.Model
      defaults:
        oid: null
        id: null
        name: null
        description: null
        current_state: null
        url: null
        tasks: []
      urlRoot: "/api/stories"

      initialize: ->
        @convertRawTasks()

      convertRawTasks: ->
        rawTasks = @get("tasks")
        if _.isArray(rawTasks)
          tasks = new Entities.Proj.Tasks(rawTasks)
          @set("tasks", tasks)

    class Entities.Proj.Stories extends CommonModel.BaseCollection
      model: Entities.Proj.Story
      url: "/api/stories"
      offset: 0

    class Entities.Proj.Project extends CommonModel.BaseModel
      defaults:
        oid: null
        id: null
        name: null
        issynced: false
        stories: []
        time_zone: null
      urlRoot: "/api/projects"
      initialize: ->
        @convertRawStories()

        # whenever the project changes (like due to fetch),
        # make sure raw stories gets converted
        @on "change", @convertRawStories

      convertRawStories: ->
        rawStories = @get("stories")
        if _.isArray(rawStories)
          stories = new Entities.Proj.Stories(rawStories)
          @set("stories", stories)

    class Entities.Proj.Projects extends CommonModel.BaseCollection
      model: Entities.Proj.Project
      url: "/api/projects"

    projects = new Entities.Proj.Projects((Projects? and Projects) or null)

    API =
      getProject: (id) ->
        # fetch project's stories and other info from backend,
        # because its not enough to just use .get from projects collection.
        defer = $.Deferred()
        project = projects.get(id)
        project.fetch
          success: (model, response, options) ->
            if response.d.redirect_to?
              Devfeed.redirect(response.d.redirect_to)
            else
              defer.resolve model
          error: (model, response, options) ->
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

      getStories: (id, filters, isReset, isRemove) ->
        defer = $.Deferred()
        project = projects.get(id)
        stories = project.get("stories")
        targetOffset = stories.size()
        stories.fetch
          reset: isReset
          remove: isRemove
          data:
            project_id: id
            offset: targetOffset
            filters: filters.join(",")
          success: (collection, response, options) ->
            if response.s == 200 and response.d.length > 0
              stories.offset = targetOffset
            defer.resolve null
          error: (collection, response, options) ->
            defer.resolve null
        return defer.promise()

    Devfeed.reqres.setHandler "project:entity", (id) ->
      return API.getProject(id)

    Devfeed.reqres.setHandler "project:entities", ->
      return API.getProjects()

    Devfeed.reqres.setHandler "project:stories", (id, filters, isReset, isRemove) ->
      return API.getStories(id, filters, isReset, isRemove)

  return Devfeed.Entities.Proj

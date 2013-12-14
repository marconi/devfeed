(function() {
  var __hasProp = {}.hasOwnProperty,
    __extends = function(child, parent) { for (var key in parent) { if (__hasProp.call(parent, key)) child[key] = parent[key]; } function ctor() { this.constructor = child; } ctor.prototype = parent.prototype; child.prototype = new ctor(); child.__super__ = parent.prototype; return child; };

  define(["devfeed", "common_model"], function(Devfeed, CommonModel) {
    Devfeed.module("Entities", function(Entities, Devfeed, Backbone, Marionette, $, _) {
      var API, projects, _ref, _ref1, _ref2, _ref3, _ref4, _ref5;
      Entities.Proj = {};
      Entities.Proj.Task = (function(_super) {
        __extends(Task, _super);

        function Task() {
          _ref = Task.__super__.constructor.apply(this, arguments);
          return _ref;
        }

        Task.prototype.defaults = {
          oid: null,
          id: null,
          position: null,
          description: null,
          complete: false,
          created_at: null,
          updated_at: null
        };

        return Task;

      })(Backbone.Model);
      Entities.Proj.Tasks = (function(_super) {
        __extends(Tasks, _super);

        function Tasks() {
          _ref1 = Tasks.__super__.constructor.apply(this, arguments);
          return _ref1;
        }

        Tasks.prototype.model = Entities.Proj.Task;

        Tasks.prototype.comparator = function(task) {
          return task.get("position");
        };

        return Tasks;

      })(Backbone.Collection);
      Entities.Proj.Story = (function(_super) {
        __extends(Story, _super);

        function Story() {
          _ref2 = Story.__super__.constructor.apply(this, arguments);
          return _ref2;
        }

        Story.prototype.defaults = {
          oid: null,
          id: null,
          name: null,
          description: null,
          current_state: null,
          url: null,
          tasks: []
        };

        Story.prototype.urlRoot = "/api/stories";

        Story.prototype.initialize = function() {
          return this.convertRawTasks();
        };

        Story.prototype.convertRawTasks = function() {
          var rawTasks, tasks;
          rawTasks = this.get("tasks");
          if (_.isArray(rawTasks)) {
            tasks = new Entities.Proj.Tasks(rawTasks);
            return this.set("tasks", tasks);
          }
        };

        return Story;

      })(Backbone.Model);
      Entities.Proj.Stories = (function(_super) {
        __extends(Stories, _super);

        function Stories() {
          _ref3 = Stories.__super__.constructor.apply(this, arguments);
          return _ref3;
        }

        Stories.prototype.model = Entities.Proj.Story;

        Stories.prototype.url = "/api/stories";

        Stories.prototype.offset = 0;

        return Stories;

      })(CommonModel.BaseCollection);
      Entities.Proj.Project = (function(_super) {
        __extends(Project, _super);

        function Project() {
          _ref4 = Project.__super__.constructor.apply(this, arguments);
          return _ref4;
        }

        Project.prototype.defaults = {
          oid: null,
          id: null,
          name: null,
          issynced: false,
          stories: [],
          time_zone: null
        };

        Project.prototype.urlRoot = "/api/projects";

        Project.prototype.initialize = function() {
          this.convertRawStories();
          return this.on("change", this.convertRawStories);
        };

        Project.prototype.convertRawStories = function() {
          var rawStories, stories;
          rawStories = this.get("stories");
          if (_.isArray(rawStories)) {
            stories = new Entities.Proj.Stories(rawStories);
            return this.set("stories", stories);
          }
        };

        return Project;

      })(CommonModel.BaseModel);
      Entities.Proj.Projects = (function(_super) {
        __extends(Projects, _super);

        function Projects() {
          _ref5 = Projects.__super__.constructor.apply(this, arguments);
          return _ref5;
        }

        Projects.prototype.model = Entities.Proj.Project;

        Projects.prototype.url = "/api/projects";

        return Projects;

      })(CommonModel.BaseCollection);
      projects = new Entities.Proj.Projects(((typeof Projects !== "undefined" && Projects !== null) && Projects) || null);
      API = {
        getProject: function(id) {
          var defer, project;
          defer = $.Deferred();
          project = projects.get(id);
          project.fetch({
            success: function(model, response, options) {
              if (response.d.redirect_to != null) {
                return Devfeed.redirect(response.d.redirect_to);
              } else {
                return defer.resolve(model);
              }
            },
            error: function(model, response, options) {
              return defer.resolve(void 0);
            }
          });
          return defer.promise();
        },
        getProjects: function() {
          var defer;
          defer = $.Deferred();
          if (projects.length === 0) {
            projects.fetch({
              success: function(collection, response, options) {
                if (response.s === 200) {
                  projects.reset(response.d);
                }
                return defer.resolve(projects);
              },
              error: function(collection, response, options) {
                return defer.resolve(null);
              }
            });
          } else {
            defer.resolve(projects);
          }
          return defer.promise();
        },
        getStories: function(id, filters, isReset, isRemove) {
          var defer, project, stories, targetOffset;
          defer = $.Deferred();
          project = projects.get(id);
          stories = project.get("stories");
          targetOffset = stories.size();
          stories.fetch({
            reset: isReset,
            remove: isRemove,
            data: {
              project_id: id,
              offset: targetOffset,
              filters: filters.join(",")
            },
            success: function(collection, response, options) {
              if (response.s === 200 && response.d.length > 0) {
                stories.offset = targetOffset;
              }
              return defer.resolve(null);
            },
            error: function(collection, response, options) {
              return defer.resolve(null);
            }
          });
          return defer.promise();
        }
      };
      Devfeed.reqres.setHandler("project:entity", function(id) {
        return API.getProject(id);
      });
      Devfeed.reqres.setHandler("project:entities", function() {
        return API.getProjects();
      });
      return Devfeed.reqres.setHandler("project:stories", function(id, filters, isReset, isRemove) {
        return API.getStories(id, filters, isReset, isRemove);
      });
    });
    return Devfeed.Entities.Proj;
  });

}).call(this);

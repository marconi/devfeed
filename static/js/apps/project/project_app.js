(function() {
  var __hasProp = {}.hasOwnProperty,
    __extends = function(child, parent) { for (var key in parent) { if (__hasProp.call(parent, key)) child[key] = parent[key]; } function ctor() { this.constructor = child; } ctor.prototype = parent.prototype; child.prototype = new ctor(); child.__super__ = parent.prototype; return child; };

  define(["devfeed", "chat", "project_list_controller", "project_show_controller"], function(Devfeed, Chat, ProjectListController, ProjectShowController) {
    Devfeed.module("ProjectApp", function(ProjectApp, Devfeed, Backbone, Marionette, $, _) {
      var API, _ref;
      ProjectApp.Router = (function(_super) {
        __extends(Router, _super);

        function Router() {
          _ref = Router.__super__.constructor.apply(this, arguments);
          return _ref;
        }

        Router.prototype.appRoutes = {
          "projects": "listProjects",
          "projects/:id": "showProject"
        };

        return Router;

      })(Marionette.AppRouter);
      API = {
        listProjects: function() {
          return Devfeed.isLoggedIn(function() {
            return ProjectListController.listProjects();
          });
        },
        showProject: function(id) {
          return Devfeed.isLoggedIn(function() {
            return ProjectShowController.showProject(id);
          });
        }
      };
      Devfeed.on("projects:list", function() {
        Devfeed.navigate("projects");
        return API.listProjects();
      });
      Devfeed.on("project:show", function(id) {
        Devfeed.navigate("projects/" + id);
        return API.showProject(id);
      });
      return Devfeed.addInitializer(function() {
        return new ProjectApp.Router({
          controller: API
        });
      });
    });
    return Devfeed.ProjectApp;
  });

}).call(this);

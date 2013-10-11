// Generated by CoffeeScript 1.6.3
(function() {
  define(["devfeed", "common_view", "project", "project_list_view"], function(Devfeed, CommonView, Proj, ProjectListView) {
    Devfeed.module("ProjectApp.List", function(List, Devfeed, Backbone, Marionette, $, _) {
      return List.Controller = {
        listProjects: function() {
          var loadingProjects, preloaderView;
          preloaderView = new CommonView.Preloader();
          Devfeed.contentRegion.show(preloaderView);
          loadingProjects = Devfeed.request("project:entities");
          return $.when(loadingProjects).done(function(projects) {
            var projectListView;
            projectListView = new ProjectListView.List({
              collection: projects
            });
            return Devfeed.contentRegion.show(projectListView);
          });
        }
      };
    });
    return Devfeed.ProjectApp.List.Controller;
  });

}).call(this);

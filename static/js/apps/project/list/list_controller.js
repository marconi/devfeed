(function() {
  define(["devfeed", "common_view", "project", "project_list_view"], function(Devfeed, CommonView, Proj, ProjectListView) {
    Devfeed.module("ProjectApp.List", function(List, Devfeed, Backbone, Marionette, $, _) {
      List.Controller = {
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
      return Devfeed.on("project:synced", function(id) {
        var loadingProjects;
        loadingProjects = Devfeed.request("project:entities");
        return $.when(loadingProjects).done(function(projects) {
          var project;
          project = projects.get(id);
          return project.set("issynced", true);
        });
      });
    });
    return Devfeed.ProjectApp.List.Controller;
  });

}).call(this);

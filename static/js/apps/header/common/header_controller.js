(function() {
  define(["devfeed", "header_view"], function(Devfeed, HeaderView) {
    Devfeed.module("HeaderApp.Common", function(Common, Devfeed, Backbone, Marionette, $, _) {
      var headerView;
      headerView = new HeaderView.Header();
      return Common.Controller = {
        showHeader: function() {
          return Devfeed.headerRegion.show(headerView);
        },
        showUserinfo: function() {
          var userSession, userinfoView;
          userSession = Devfeed.request("user:session");
          userinfoView = new HeaderView.UserinfoView({
            model: userSession
          });
          return headerView.userinfoRegion.show(userinfoView);
        },
        removeUserinfo: function() {
          return headerView.userinfoRegion.close();
        }
      };
    });
    return Devfeed.HeaderApp.Common.Controller;
  });

}).call(this);

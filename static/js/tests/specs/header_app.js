(function() {
  define(["devfeed", "header_controller", "header_app"], function(Devfeed, HeaderController, HeaderApp) {
    var userSessionHandler;
    userSessionHandler = null;
    before(function() {
      userSessionHandler = Devfeed.reqres.getHandler("user:session");
      Devfeed.reqres.removeHandler("user:session");
      return Devfeed.reqres.setHandler("user:session", function() {
        var userSession;
        userSession = new Devfeed.Entities.UserSession({
          sessionId: 1234,
          name: "Marconi Moreto"
        });
        return userSession;
      });
    });
    after(function() {
      Devfeed.reqres.removeHandler("user:session");
      return Devfeed.reqres.setHandler("user:session", userSessionHandler);
    });
    describe("HeaderApp", function() {
      return it("shows header on Devfeed start", function() {
        Backbone.history.constructor.started = false;
        expect($('#header-region #header')[0]).to.not.exist;
        Devfeed.start();
        expect($('#header-region #header')[0]).to.exist;
        return Devfeed.headerRegion.close();
      });
    });
    return describe("HeaderApp.Common.Controller", function() {
      it("showHeader displays header", function() {
        expect($('#header-region #header')[0]).to.not.exist;
        HeaderController.showHeader();
        return expect($('#header-region #header')[0]).to.exist;
      });
      return it("showUserinfo displays logged-in user info", function() {
        expect($('#header-region #userinfo')[0]).to.not.exist;
        HeaderController.showUserinfo();
        expect($('#header-region #userinfo')[0]).to.exist;
        return expect($("#user-name").html()).to.equal("Marconi Moreto");
      });
    });
  });

}).call(this);

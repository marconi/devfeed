(function() {
  define(["devfeed", "login_view", "login_controller", "login_controller", "auth_app"], function(Devfeed, LoginView, LoginController, LogoutController, AuthApp) {
    before(function() {
      return Backbone.history.start({
        pushState: true
      });
    });
    afterEach(function() {
      return Devfeed.contentRegion.close();
    });
    after(function() {
      return Backbone.history.stop();
    });
    describe("AuthApp", function() {
      return it("request auth:login:show displays login form and url", function() {
        expect($('#content-region #login')[0]).to.not.exist;
        Devfeed.trigger("auth:login:show");
        expect($('#content-region #login')[0]).to.exist;
        return expect(window.location.pathname).to.equal("/login");
      });
    });
    describe("AuthApp.Login.View", function() {
      it("displays errors on empty form", function() {
        var loginView;
        loginView = new LoginView.LoginForm();
        loginView.on("auth:login", function(data) {
          return loginView.triggerMethod("form:data:invalid", data);
        });
        Devfeed.contentRegion.show(loginView);
        $("#login .signin").click();
        return expect($("#login span.error").length).to.equal(2);
      });
      return it("displays alert on invalid credentials", function() {
        var loginView;
        loginView = new LoginView.LoginForm();
        loginView.on("auth:login", function(data) {
          return loginView.triggerMethod("form:data:invalid", data);
        });
        Devfeed.contentRegion.show(loginView);
        loginView.ui.emailInput.val("foo@bar.com");
        loginView.ui.passwordInput.val("secret");
        $("#login .signin").click();
        return expect($("#login-alert-region .alert-box")).to.exist;
      });
    });
    return describe("AuthApp.Login.Controller", function() {
      return it("showLogin displays login form", function() {
        expect($('#content-region #login')[0]).to.not.exist;
        LoginController.showLogin();
        return expect($('#content-region #login')[0]).to.exist;
      });
    });
  });

}).call(this);

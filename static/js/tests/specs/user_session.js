(function() {
  define(["devfeed", "user_session"], function(Devfeed, UserSession) {
    return describe("Entities.UserSession", function() {
      it("returns false when no session", sinon.test(function() {
        var userSession;
        this.stub(UserSession.prototype, "isAuthenticated").returns(false);
        userSession = new UserSession();
        return expect(userSession.isAuthenticated()).to.be["false"];
      }));
      it("returns true when there's session", sinon.test(function() {
        var userSession;
        userSession = new UserSession({
          sessionId: 1234
        });
        return expect(userSession.isAuthenticated()).to.be["true"];
      }));
      it("request session:authenticated returns false when not authenticated", sinon.test(function() {
        this.stub(UserSession.prototype, "isAuthenticated").returns(false);
        return expect(Devfeed.request("session:authenticated")).to.be["false"];
      }));
      it("request session:authenticated returns true when authenticated", sinon.test(function() {
        this.stub(UserSession.prototype, "isAuthenticated").returns(true);
        return expect(Devfeed.request("session:authenticated")).to.be["true"];
      }));
      return it("request user:session returns UserSession instance", function() {
        return expect(Devfeed.request("user:session")).to.be["instanceof"](UserSession);
      });
    });
  });

}).call(this);

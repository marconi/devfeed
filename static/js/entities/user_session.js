(function() {
  var __hasProp = {}.hasOwnProperty,
    __extends = function(child, parent) { for (var key in parent) { if (__hasProp.call(parent, key)) child[key] = parent[key]; } function ctor() { this.constructor = child; } ctor.prototype = parent.prototype; child.prototype = new ctor(); child.__super__ = parent.prototype; return child; };

  define(["devfeed", "validation"], function(Devfeed, Validation) {
    Devfeed.module("Entities", function(Entities, Devfeed, Backbone, Marionette, $, _) {
      var API, userSession, _ref;
      Entities.UserSession = (function(_super) {
        __extends(UserSession, _super);

        function UserSession() {
          _ref = UserSession.__super__.constructor.apply(this, arguments);
          return _ref;
        }

        UserSession.prototype.defaults = {
          sessionId: null,
          name: null,
          email: null,
          password: null,
          apitoken: null
        };

        UserSession.prototype.validateName = false;

        UserSession.prototype.validateEmail = true;

        UserSession.prototype.validatePassword = true;

        UserSession.prototype.validateApiToken = false;

        UserSession.prototype.validate = function(attrs, options) {
          var errors;
          errors = {};
          if (this.validateName && !attrs.name) {
            errors.name = "Name is required";
          }
          if (this.validateEmail && !attrs.email) {
            errors.email = "Email is required";
          } else if (this.validateEmail && !Validation.patterns.email.test(attrs.email)) {
            errors.email = "Must be a valid email";
          }
          if (this.validatePassword && !attrs.password) {
            errors.password = "Password is required";
          }
          if (this.validateApiToken && !attrs.apitoken) {
            errors.apitoken = "API Token is required";
          }
          if (!_.isEmpty(errors)) {
            return errors;
          }
        };

        UserSession.prototype.resetValidationRules = function() {
          this.set({
            name: null,
            email: null,
            password: null,
            apitoken: null
          });
          this.validateName = false;
          this.validateEmail = true;
          this.validatePassword = true;
          return this.validateApiToken = false;
        };

        UserSession.prototype.backupAttrs = function() {
          return this.attrsBackup = this.attributes;
        };

        UserSession.prototype.restoreAttrs = function() {
          return this.set(this.attrsBackup);
        };

        UserSession.prototype.isLoggedIn = function() {
          var defer,
            _this = this;
          defer = $.Deferred();
          if (!this.get('sessionId')) {
            $.ajax({
              async: false,
              type: 'GET',
              url: "/isloggedin",
              dataType: "json",
              statusCode: {
                200: function(data, textStatus, xhr) {
                  _this.set("sessionId", data.sessionid);
                  _this.set("id", data.id);
                  _this.set("name", data.name);
                  _this.set("email", data.email);
                  return _this.set("apitoken", data.apitoken);
                }
              },
              complete: function() {
                var isLoggedIn;
                isLoggedIn = Boolean(_this.get('sessionId'));
                if (isLoggedIn) {
                  Devfeed.trigger("loggedin");
                }
                return defer.resolve(isLoggedIn);
              }
            });
          } else {
            defer.resolve(Boolean(this.get('sessionId')));
          }
          return defer.promise();
        };

        UserSession.prototype.login = function(email, password) {
          var defer,
            _this = this;
          defer = $.Deferred();
          $.ajax({
            async: false,
            type: "POST",
            url: "/login",
            dataType: "json",
            data: {
              email: email,
              password: password
            },
            statusCode: {
              200: function(data, textStatus, xhr) {
                _this.set("sessionId", data.sessionid);
                _this.set("id", data.id);
                _this.set("name", data.name);
                _this.set("email", data.email);
                _this.set("apitoken", data.apitoken);
                Devfeed.trigger("loggedin");
                return defer.resolve(null);
              },
              401: function(xhr, textStatus, error) {
                return defer.resolve({
                  errors: {},
                  message: xhr.responseJSON
                });
              },
              500: function() {
                return defer.resolve({
                  errors: {},
                  message: {
                    body: Devfeed.messages[500],
                    type: "alert"
                  }
                });
              }
            },
            complete: function() {
              return _this.set({
                password: null
              });
            }
          });
          return defer.promise();
        };

        UserSession.prototype.register = function(name, email, password) {
          var defer,
            _this = this;
          defer = $.Deferred();
          $.ajax({
            type: "POST",
            url: "/register",
            dataType: "json",
            data: {
              name: name,
              email: email,
              password: password
            },
            statusCode: {
              200: function(data, textStatus, xhr) {
                return defer.resolve({
                  message: xhr.responseJSON
                });
              },
              400: function(xhr, textStatus, error) {
                return defer.resolve({
                  errors: xhr.responseJSON
                });
              },
              500: function() {
                return defer.resolve({
                  errors: {},
                  message: {
                    body: Devfeed.messages[500],
                    type: "alert"
                  }
                });
              }
            },
            complete: function() {
              return _this.set({
                name: null,
                email: null,
                password: null
              });
            }
          });
          return defer.promise();
        };

        UserSession.prototype.resendActivation = function(email) {
          var defer,
            _this = this;
          defer = $.Deferred();
          $.ajax({
            type: "POST",
            url: "/activation/resend",
            dataType: "json",
            data: {
              email: email
            },
            statusCode: {
              200: function(data, textStatus, xhr) {
                return defer.resolve({
                  message: xhr.responseJSON
                });
              },
              400: function(xhr, textStatus, error) {
                return defer.resolve({
                  errors: xhr.responseJSON
                });
              },
              500: function() {
                return defer.resolve({
                  errors: {},
                  message: {
                    body: Devfeed.messages[500],
                    type: "alert"
                  }
                });
              }
            },
            complete: function() {
              return _this.set({
                email: null
              });
            }
          });
          return defer.promise();
        };

        UserSession.prototype.retrievePassword = function(email) {
          var defer,
            _this = this;
          defer = $.Deferred();
          $.ajax({
            type: "POST",
            url: "/password/retrieve",
            dataType: "json",
            data: {
              email: email
            },
            statusCode: {
              200: function(data, textStatus, xhr) {
                return defer.resolve({
                  message: xhr.responseJSON
                });
              },
              400: function(xhr, textStatus, error) {
                return defer.resolve({
                  errors: xhr.responseJSON
                });
              },
              500: function() {
                return defer.resolve({
                  errors: {},
                  message: {
                    body: Devfeed.messages[500],
                    type: "alert"
                  }
                });
              }
            },
            complete: function() {
              return _this.set({
                email: null
              });
            }
          });
          return defer.promise();
        };

        UserSession.prototype.updateSettings = function(fields) {
          var defer,
            _this = this;
          fields = _.compactObject(fields);
          defer = $.Deferred();
          $.ajax({
            type: "POST",
            url: "/settings/update",
            dataType: "json",
            data: fields,
            statusCode: {
              200: function(data, textStatus, xhr) {
                _this.set(data.user);
                Devfeed.execute("show:userinfo");
                return defer.resolve({
                  message: data.message
                });
              },
              400: function(xhr, textStatus, error) {
                return defer.resolve({
                  errors: xhr.responseJSON
                });
              },
              500: function() {
                return defer.resolve({
                  errors: {},
                  message: {
                    body: Devfeed.messages[500],
                    type: "alert"
                  }
                });
              }
            }
          });
          return defer.promise();
        };

        return UserSession;

      })(Backbone.Model);
      userSession = new Devfeed.Entities.UserSession();
      API = {
        isLoggedIn: function() {
          return userSession.isLoggedIn();
        },
        getUserSession: function() {
          return userSession;
        }
      };
      Devfeed.reqres.setHandler("session:isloggedin", function() {
        return API.isLoggedIn();
      });
      Devfeed.reqres.setHandler("session:login", function(email, password) {
        var defer;
        userSession.set({
          email: email,
          password: password
        });
        if (!userSession.isValid()) {
          defer = $.Deferred();
          defer.resolve({
            errors: userSession.validationError
          });
          userSession.resetValidationRules();
          return defer.promise();
        }
        userSession.resetValidationRules();
        return userSession.login(email, password);
      });
      Devfeed.reqres.setHandler("session:register", function(name, email, password) {
        var defer;
        userSession.validateName = true;
        userSession.set({
          name: name,
          email: email,
          password: password
        });
        if (!userSession.isValid()) {
          userSession.resetValidationRules();
          defer = $.Deferred();
          defer.resolve({
            errors: userSession.validationError
          });
          return defer.promise();
        }
        userSession.resetValidationRules();
        return userSession.register(name, email, password);
      });
      Devfeed.reqres.setHandler("session:activation:resend", function(email) {
        var defer;
        userSession.validatePassword = false;
        userSession.set({
          email: email
        });
        if (!userSession.isValid()) {
          userSession.resetValidationRules();
          defer = $.Deferred();
          defer.resolve({
            errors: userSession.validationError
          });
          return defer.promise();
        }
        userSession.resetValidationRules();
        return userSession.resendActivation(email);
      });
      Devfeed.reqres.setHandler("session:password:retrieve", function(email) {
        var defer;
        userSession.validatePassword = false;
        userSession.set({
          email: email
        });
        if (!userSession.isValid()) {
          userSession.resetValidationRules();
          defer = $.Deferred();
          defer.resolve({
            errors: userSession.validationError
          });
          return defer.promise();
        }
        userSession.resetValidationRules();
        return userSession.retrievePassword(email);
      });
      Devfeed.reqres.setHandler("settings:update:general", function(form) {
        var defer;
        userSession.backupAttrs();
        userSession.validateName = true;
        userSession.validatePassword = false;
        userSession.set(fields);
        if (!userSession.isValid()) {
          userSession.restoreAttrs();
          defer = $.Deferred();
          defer.resolve({
            errors: userSession.validationError
          });
          return defer.promise();
        }
        userSession.resetValidationRules();
        return userSession.updateSettings(fields);
      });
      Devfeed.reqres.setHandler("settings:update:pivotal", function(fields) {
        var defer;
        userSession.backupAttrs();
        userSession.validateEmail = false;
        userSession.validateApiToken = true;
        userSession.validatePassword = false;
        userSession.set(fields);
        if (!userSession.isValid()) {
          userSession.restoreAttrs();
          defer = $.Deferred();
          defer.resolve({
            errors: userSession.validationError
          });
          return defer.promise();
        }
        userSession.resetValidationRules();
        return userSession.updateSettings(fields);
      });
      return Devfeed.reqres.setHandler("user:session", function() {
        return API.getUserSession();
      });
    });
    return Devfeed.Entities.UserSession;
  });

}).call(this);

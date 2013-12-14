define ["devfeed", "validation"], (Devfeed, Validation) ->

  Devfeed.module "Entities", (Entities, Devfeed, Backbone, Marionette, $, _) ->

    class Entities.UserSession extends Backbone.Model
      defaults:
        sessionId: null
        name: null
        email: null
        password: null
        apitoken: null
      validateName: false # flag to toggle name validation
      validateEmail: true # flag to toggle email validation
      validatePassword: true # flag to toggle password validation
      validateApiToken: false # flag to toggle token validation

      validate: (attrs, options) ->
        errors = {}
        if @validateName and not attrs.name
          errors.name = "Name is required"
        if @validateEmail and not attrs.email
          errors.email = "Email is required"
        else if @validateEmail and not Validation.patterns.email.test(attrs.email)
          errors.email = "Must be a valid email"
        if @validatePassword and not attrs.password
          errors.password = "Password is required"
        if @validateApiToken and not attrs.apitoken
          errors.apitoken = "API Token is required"
        if not _.isEmpty(errors)
          return errors

      resetValidationRules: ->
        @set name: null, email: null, password: null, apitoken: null
        @validateName = false
        @validateEmail = true
        @validatePassword = true
        @validateApiToken = false

      backupAttrs: ->
        @attrsBackup = @attributes

      restoreAttrs: ->
        @set @attrsBackup

      isLoggedIn: ->
        defer = $.Deferred()
        # if we don't have session id, fetch from backend
        if not @get('sessionId')
          $.ajax
            async: false
            type: 'GET'
            url: "/isloggedin"
            dataType: "json"
            statusCode:
              200: (data, textStatus, xhr) =>
                @set("sessionId", data.sessionid)
                @set("id", data.id)
                @set("name", data.name)
                @set("email", data.email)
                @set("apitoken", data.apitoken)
            complete: =>
              isLoggedIn = Boolean(@get('sessionId'))
              if isLoggedIn
                # triger loggedin event so any listener like
                # creating websocket can react
                Devfeed.trigger("loggedin")
              defer.resolve isLoggedIn
        else
          defer.resolve Boolean(@get('sessionId'))
        return defer.promise()

      login: (email, password) ->
        defer = $.Deferred()
        $.ajax
          async: false
          type: "POST"
          url: "/login"
          dataType: "json"
          data:
            email: email
            password: password
          statusCode:
            200: (data, textStatus, xhr) =>
              @set("sessionId", data.sessionid)
              @set("id", data.id)
              @set("name", data.name)
              @set("email", data.email)
              @set("apitoken", data.apitoken)
              Devfeed.trigger("loggedin")
              defer.resolve null
            401: (xhr, textStatus, error) =>
              defer.resolve
                errors: {}
                message: xhr.responseJSON
            500: =>
              defer.resolve
                errors: {}
                message:
                  body: Devfeed.messages[500], type: "alert"
          complete: =>
            @set password: null
        return defer.promise()

      register: (name, email, password) ->
        defer = $.Deferred()
        $.ajax
          type: "POST"
          url: "/register"
          dataType: "json"
          data:
            name: name
            email: email
            password: password
          statusCode:
            200: (data, textStatus, xhr) =>
              defer.resolve message: xhr.responseJSON
            400: (xhr, textStatus, error) =>
              defer.resolve errors: xhr.responseJSON
            500: =>
              defer.resolve
                errors: {}
                message:
                  body: Devfeed.messages[500], type: "alert"
          complete: =>
            @set name: null, email: null, password: null
        return defer.promise()

      resendActivation: (email) ->
        defer = $.Deferred()
        $.ajax
          type: "POST"
          url: "/activation/resend"
          dataType: "json"
          data:
            email: email
          statusCode:
            200: (data, textStatus, xhr) =>
              defer.resolve message: xhr.responseJSON
            400: (xhr, textStatus, error) =>
              defer.resolve errors: xhr.responseJSON
            500: =>
              defer.resolve
                errors: {}
                message:
                  body: Devfeed.messages[500], type: "alert"
          complete: =>
            @set email: null
        return defer.promise()

      retrievePassword: (email) ->
        defer = $.Deferred()
        $.ajax
          type: "POST"
          url: "/password/retrieve"
          dataType: "json"
          data:
            email: email
          statusCode:
            200: (data, textStatus, xhr) =>
              defer.resolve message: xhr.responseJSON
            400: (xhr, textStatus, error) =>
              defer.resolve errors: xhr.responseJSON
            500: =>
              defer.resolve
                errors: {}
                message:
                  body: Devfeed.messages[500], type: "alert"
          complete: =>
            @set email: null
        return defer.promise()

      updateSettings: (fields) ->
        fields = _.compactObject(fields)
        defer = $.Deferred()
        $.ajax
          type: "POST"
          url: "/settings/update"
          dataType: "json"
          data: fields # only update present fields
          statusCode:
            200: (data, textStatus, xhr) =>
              @set data.user
              Devfeed.execute("show:userinfo")
              defer.resolve message: data.message
            400: (xhr, textStatus, error) =>
              defer.resolve errors: xhr.responseJSON
            500: =>
              defer.resolve
                errors: {}
                message:
                  body: Devfeed.messages[500], type: "alert"
        return defer.promise()

    # create new user session
    userSession = new Devfeed.Entities.UserSession()

    API =
      isLoggedIn: ->
        return userSession.isLoggedIn()

      getUserSession: ->
        return userSession

    Devfeed.reqres.setHandler "session:isloggedin", ->
      return API.isLoggedIn()

    Devfeed.reqres.setHandler "session:login", (email, password) ->
      userSession.set email: email, password: password
      if not userSession.isValid()
        defer = $.Deferred()
        defer.resolve errors: userSession.validationError
        userSession.resetValidationRules()
        return defer.promise()
      userSession.resetValidationRules()
      return userSession.login(email, password)

    Devfeed.reqres.setHandler "session:register", (name, email, password) ->
      userSession.validateName = true
      userSession.set name: name, email: email, password: password
      if not userSession.isValid()
        userSession.resetValidationRules()
        defer = $.Deferred()
        defer.resolve errors: userSession.validationError
        return defer.promise()
      userSession.resetValidationRules()
      return userSession.register(name, email, password)

    Devfeed.reqres.setHandler "session:activation:resend", (email) ->
      userSession.validatePassword = false
      userSession.set email: email
      if not userSession.isValid()
        userSession.resetValidationRules()
        defer = $.Deferred()
        defer.resolve errors: userSession.validationError
        return defer.promise()
      userSession.resetValidationRules()
      return userSession.resendActivation(email)

    Devfeed.reqres.setHandler "session:password:retrieve", (email) ->
      userSession.validatePassword = false
      userSession.set email: email
      if not userSession.isValid()
        userSession.resetValidationRules()
        defer = $.Deferred()
        defer.resolve errors: userSession.validationError
        return defer.promise()
      userSession.resetValidationRules()
      return userSession.retrievePassword(email)

    Devfeed.reqres.setHandler "settings:update:general", (form) ->
      userSession.backupAttrs()

      # make sure name is validated when updating general settings
      userSession.validateName = true

      # optionally validate password since it can be omitted when just
      # updating the other fields.
      userSession.validatePassword = false

      userSession.set fields
      if not userSession.isValid()
        userSession.restoreAttrs()
        defer = $.Deferred()
        defer.resolve errors: userSession.validationError
        return defer.promise()
      userSession.resetValidationRules()
      return userSession.updateSettings(fields)

    Devfeed.reqres.setHandler "settings:update:pivotal", (fields) ->
      userSession.backupAttrs()

      # make sure api token is validated when updating pivotal settings
      userSession.validateEmail = false
      userSession.validateApiToken = true

      # pivotal has no password settings
      userSession.validatePassword = false

      userSession.set fields
      if not userSession.isValid()
        userSession.restoreAttrs()
        defer = $.Deferred()
        defer.resolve errors: userSession.validationError
        return defer.promise()
      userSession.resetValidationRules()
      return userSession.updateSettings(fields)

    Devfeed.reqres.setHandler "user:session", ->
      return API.getUserSession()

  return Devfeed.Entities.UserSession

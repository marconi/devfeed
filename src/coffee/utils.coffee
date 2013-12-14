_.mixin
  compactObject: (o) ->
    _.each o, (v, k) ->
      if not v
        delete o[k]
    return o

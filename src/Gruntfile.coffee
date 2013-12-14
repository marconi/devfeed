module.exports = (grunt) ->

  grunt.initConfig
    pkg: grunt.file.readJSON('package.json')
    compass:
      dist:
        options:
          require: 'zurb-foundation'
          outputStyle: 'compressed'
          sassDir: 'sass'
          cssDir: '../static/css'
          imagesPath: '../static/images'
          environment: 'production'
    coffee:
      glob_to_multiple:
        expand: true
        cwd: 'coffee/'
        src: ['**/*.coffee']
        dest: '../static/js'
        ext: '.js'
    requirejs:
      compile:
        options:
          baseUrl: '../static/js'
          name: 'vendors/almond'
          include: 'main'
          mainConfigFile: '../static/js/main.js'
          insertRequire: ['main']
          out: '../static/js/main.built.js'
          findNestedDependencies: true
    watch:
      coffee:
        files: ['**/*.coffee']
        tasks: 'coffee'
      compass:
        files: ['**/*.scss']
        tasks: 'compass'

  grunt.loadNpmTasks('grunt-contrib-compass')
  grunt.loadNpmTasks('grunt-contrib-coffee')
  grunt.loadNpmTasks('grunt-contrib-requirejs')
  grunt.loadNpmTasks('grunt-contrib-watch')

  grunt.registerTask('default', ['compass', 'coffee', 'requirejs'])

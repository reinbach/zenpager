var gulp = require('gulp'),
    vulcanize = require("gulp-vulcanize");
var $ = require("gulp-load-plugins")();
var dest_dir = __dirname + "/dist/";

gulp.task("vulcanize", function() {
    return gulp.src("src/index.html")
        .pipe(vulcanize({
            dest: dest_dir,
            inline: true,
            csp: true,
            strip: true
        }))
        .pipe($.replace(/\.\.\/src\/fonts/g, 'fonts'))
        .pipe(gulp.dest(dest_dir))
        .pipe(gulp.src("*.js")
              .pipe($.uglify())
              .on('error', function (error) {
                  console.error('' + error);
              })
              .pipe($.notify({message: 'Vulcanize task completed'})));
});

gulp.task('css', function() {
    return gulp.src("src/scss/*.scss")
        .pipe($.autoprefixer({browsers: '> 1%', cascade: false, remove: true}))
        .pipe($.concat('app.css'))
        .pipe($.sass({errLogToConsole: true}))
        .pipe($.minifyCss())
        .pipe($.replace(/\.\.\/fonts/g, 'fonts'))
        .pipe(gulp.dest(__dirname + "/src/"))
        .on('error', function (error) {
            console.error('' + error);
        })
        .pipe($.notify({message: 'Style task completed'}));
});

gulp.task('fonts', function() {
    return gulp.src("src/fonts/*")
        .pipe(gulp.dest(dest_dir + "/fonts"));
});

gulp.task('build', $.sequence('css', ['vulcanize', 'fonts']));

gulp.task('default', function() {
    gulp.watch("src/scss/*", ['css']);
    // gulp.watch("src/*", ['build']);
});


// Workaround for https://github.com/gulpjs/gulp/issues/71
var origSrc = gulp.src;
gulp.src = function () {
    return fixPipe(origSrc.apply(this, arguments));
};
function fixPipe(stream) {
    var origPipe = stream.pipe;
    stream.pipe = function (dest) {
        arguments[0] = dest.on('error', function (error) {
            var nextStreams = dest._nextStreams;
            if (nextStreams) {
                nextStreams.forEach(function (nextStream) {
                    nextStream.emit('error', error);
                });
            } else if (dest.listeners('error').length === 1) {
                throw error;
            }
        });
        var nextStream = fixPipe(origPipe.apply(this, arguments));
        (this._nextStreams || (this._nextStreams = [])).push(nextStream);
        return nextStream;
    };
    return stream;
}

gulp.task('list', function() {
    return gulp.src(jsFiles.vendor)
        .pipe(gulp.dest("/tmp/"));
});

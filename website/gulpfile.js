var gulp = require('gulp'),
    vulcanize = require("gulp-vulcanize");
var $ = require("gulp-load-plugins")();
var project_dir = "website/";
var dest_dir = __dirname;

function processFiles(files, name) {
    return gulp.src(files)
        .pipe($.rename(name + ".html"))
        .pipe(vulcanize({
            dest: dest_dir,
            inline: true,
            csp: true,
            strip: true
        }))
        .pipe(gulp.dest(dest_dir))
        .pipe(gulp.src("*.js")
              .pipe($.uglify())
              .on('error', function (error) {
                  console.error('' + error);
              })
              .pipe($.notify({message: name + ' task complete'})));
}

gulp.task('intro', function() {
    processFiles("intro/index.html", 'intro');
});

gulp.task('dashboard', function() {
    processFiles("dashboard/index.html", 'dashboard');
});

// grouped tasks
gulp.task('build', function() {
    gulp.start('intro', 'dashboard');
});

// main
gulp.task('default', function() {
    // watch intro
    gulp.watch("intro/*", ['intro']);
    gulp.watch("dashboard/*", ['dashboard']);
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

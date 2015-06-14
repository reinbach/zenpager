var gulp = require('gulp'),
    bowerFiles = require("main-bower-files");
var $ = require("gulp-load-plugins")();
var projectDir = "template/static/";
var jsFiles = {
    "vendor": bowerFiles({filter: '**/*.js'}),
    "local": {
        "intro": projectDir + "src/js/intro/*.jsx",
        "dashboard": prependPath(projectDir + "src/js/dashboard/", [
            "utils.jsx",
            "_auth.jsx",
            "_dashboard.jsx",
            "_profile.jsx",
            "_settings.jsx",
            "_settings_contacts.jsx",
            "_validation.jsx",
            "main.jsx"
        ])
    }
};
var cssFiles = {
    "vendor": bowerFiles({filter: '**/*.css'}),
    "local": {
        "intro": projectDir + "src/css/intro/*.scss",
        "dashboard": projectDir + "src/css/dashboard/*.scss"
    }
};

function prependPath(path, arr) {
    for (var i = 0; i < arr.length; i++) {
        arr[i] = path + arr[i];
    }
    return arr;
}

function processJSFiles(files, name) {
    return gulp.src(files)
        .pipe($.sourcemaps.init())
        .pipe($.react({harmony: true}))
        .pipe($.concat(name + '.js'))
        .pipe($.rename({ suffix: '.min' }))
        .pipe($.uglify())
        .pipe($.sourcemaps.write('.'))
        .pipe(gulp.dest(projectDir + 'dist/js'))
        .on('error', function (error) {
            console.error('' + error);
        })
        .pipe($.notify({message: 'Javascript local ' + name + ' task complete'}));
}

// js
gulp.task('js-vendor', function() {
    return gulp.src(jsFiles.vendor)
        .pipe($.concat('vendor.js'))
        .pipe($.rename({ suffix: '.min' }))
        .pipe($.uglify())
        .pipe(gulp.dest(projectDir + 'dist/js'))
        .on('error', function (error) {
            console.error('' + error);
        })
        .pipe($.notify({message: 'Javascript vendor task complete'}));
});

gulp.task('js-local-intro', function() {
    processJSFiles(jsFiles.local.intro, 'intro');
});

gulp.task('js-local-dashboard', function() {
    processJSFiles(jsFiles.local.dashboard, 'dashboard');
});

gulp.task('js-local', function() {
    gulp.start('js-local-intro', 'js-local-dashboard');
});

gulp.task('js', function() {
    gulp.start('js-vendor', 'js-local');
});

// css
gulp.task('css-vendor', function() {
    return gulp.src(cssFiles.vendor)
        .pipe($.concat('vendor.css'))
        .pipe($.rename({ suffix: '.min' }))
        .pipe($.minifyCss())
        .pipe(gulp.dest(projectDir + 'dist/css'))
        .on('error', function (error) {
            console.error('' + error);
        })
        .pipe($.notify({message: 'Style vendor task complete'}));
});

function processCSSFiles(files, name) {
    return gulp.src(files)
        .pipe($.sourcemaps.init())
        .pipe($.autoprefixer({browsers: '> 1%', cascade: false, remove: true}))
        .pipe($.concat(name + '.css'))
        .pipe($.rename({ suffix: '.min' }))
        .pipe($.sass({errLogToConsole: true}))
        .pipe($.sourcemaps.write('.'))
        .pipe($.minifyCss())
        .pipe(gulp.dest(projectDir + 'dist/css'))
        .on('error', function (error) {
            console.error('' + error);
        })
        .pipe($.notify({message: 'Style local ' + name + ' task complete'}));
}

gulp.task('css-local-intro', function() {
    processCSSFiles(cssFiles.local.intro, 'intro');
});

gulp.task('css-local-dashboard', function() {
    processCSSFiles(cssFiles.local.dashboard, 'dashboard');
});

gulp.task('css-local', function() {
    gulp.start('css-local-intro', 'css-local-dashboard');
});

gulp.task('css', function() {
    gulp.start('css-vendor', 'css-local');
});

// images
//TODO maybe compress images (imagemin)?

// fonts
gulp.task('fonts-vendor', function() {
    return gulp.src(bowerFiles({filter: '**/*.{eot,svg,ttf,woff,woff2}'}))
        .pipe(gulp.dest(projectDir + 'dist/fonts'))
        .on('error', function (error) {
            console.error('' + error);
        })
        .pipe($.notify({message: 'Fonts vendor file written'}));
});

gulp.task('fonts', function() {
    gulp.start('fonts-vendor');
});

// grouped tasks
gulp.task('local', function() {
    gulp.start('css-local', 'js-local');
});

gulp.task('vendor', function() {
    gulp.start('css-vendor', 'js-vendor', 'fonts-vendor');
});

gulp.task('build', function() {
    gulp.start('local', 'vendor', 'fonts');
});

// main
gulp.task('default', function() {
    // watch css
    gulp.watch(cssFiles.local.intro, ['css-local-intro']);
    gulp.watch(cssFiles.local.dashboard, ['css-local-dashboard']);

    // watch js
    gulp.watch(jsFiles.local.intro, ['js-local-intro']);
    gulp.watch(jsFiles.local.dashboard, ['js-local-dashboard']);
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

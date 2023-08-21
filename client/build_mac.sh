# convert the icons
# SystrayApp.app/
#   Contents/
#     Info.plist
#     MacOS/
#       go-executable
#     Resources/
#       SystrayApp.icns

# remove the old builds
rm -rf build
# set the app name as variable
APPNAME=LanCommander

# create the build folder
mkdir build
# create the app folder
mkdir build/$APPNAME.app
# create the contents folder
mkdir build/$APPNAME.app/Contents
# create the MacOS folder
mkdir build/$APPNAME.app/Contents/MacOS
# create the Resources folder
mkdir build/$APPNAME.app/Contents/Resources


# copy the Info.plist
cp Info.plist build/$APPNAME.app/Contents/

# copy all files from assets folder to MacOS/assets folder
# cp -r assets build/$APPNAME.app/Contents/MacOS/

# convert the icon
icnsify -i assets/link-green.png -o build/$APPNAME.app/Contents/Resources/icon.icns

# build the app
go build -o build/$APPNAME.app/Contents/MacOS/$APPNAME




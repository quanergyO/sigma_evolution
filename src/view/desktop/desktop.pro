TEMPLATE = app

TARGET = SigmaEvolution.out

QT       += core gui

greaterThan(QT_MAJOR_VERSION, 4): QT += widgets

CONFIG += c++20

# You can make your code fail to compile if it uses deprecated APIs.
# In order to do so, uncomment the following line.
#DEFINES += QT_DISABLE_DEPRECATED_BEFORE=0x060000    # disables all the APIs deprecated before Qt 6.0.0

BASE_DIR = $$PWD/../../..
INCLUDE_DIR = $$BASE_DIR/include
SRC_DIR = $$BASE_DIR/src
UI_DIR = $$BASE_DIR/ui

INCLUDEPATH += $$INCLUDE_DIR

SOURCES += \
    main.cpp \
    mainwindow.cpp \
    $$SRC_DIR/model/model.cpp \
    $$SRC_DIR/controller/controller.cpp

HEADERS += \
    $$INCLUDE_DIR/view/desktop/mainwindow.h \
    $$INCLUDE_DIR/model/model.h \
    $$INCLUDE_DIR/controller/controller.h

FORMS += \
    $$UI_DIR/mainwindow.ui

# Default rules for deployment.
qnx: target.path = /tmp/$${TARGET}/bin
else: unix:!android: target.path = /opt/$${TARGET}/bin
!isEmpty(target.path): INSTALLS += target

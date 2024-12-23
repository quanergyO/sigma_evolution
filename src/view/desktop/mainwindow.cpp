#include "view/desktop/mainwindow.h"
#include "ui_mainwindow.h"

MainWindow::MainWindow(controller::Controller& controller, QWidget *parent)
    : QMainWindow(parent)
    , ui(new Ui::MainWindow)
    , controller_(controller)
{
    ui->setupUi(this);
}

MainWindow::~MainWindow()
{
    delete ui;
}

#ifndef MAINWINDOW_H
#define MAINWINDOW_H

#include <QMainWindow>
#include "controller/controller.h"

QT_BEGIN_NAMESPACE
namespace Ui {
class MainWindow;
}
QT_END_NAMESPACE

class MainWindow : public QMainWindow
{
    Q_OBJECT

public:
    MainWindow(controller::Controller& controller, QWidget *parent = nullptr);
    ~MainWindow();

private:
    Ui::MainWindow *ui;
    controller::Controller& controller_;
};
#endif // MAINWINDOW_H

#include "view/desktop/mainwindow.h"

#include <QApplication>

int main(int argc, char *argv[])
{
    model::Model model;
    controller::Controller controller(model);

    QApplication application(argc, argv);
    MainWindow window(controller);
    window.show();
    return application.exec();
}

#ifndef CONTROLLER_CONTROLLER_H
#define CONTROLLER_CONTROLLER_H

#include "model/model.h"

namespace controller
{

class Controller
{
public:
    Controller(model::Model &model);

private:
    model::Model& model_;
};

} // namespace controller

#endif // CONTROLLER_CONTROLLER_H

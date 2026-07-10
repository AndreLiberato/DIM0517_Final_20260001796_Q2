#include <cassert>
#include <iostream>

#include "Calc.h"

int main() {
    assert(add(3, 4) == 7);
    assert(sub(4, 3) == 1);
    assert(sub(3, 4) == -1);

    std::cout << "All tests passed." << std::endl;
    return 0;
}

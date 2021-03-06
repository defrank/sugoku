=======
Sugoku_
=======

.. _Sugoku: https://github.com/defrank/sugoku

:Authors:
    Derek M Frank

Sudoku written in Go!

In an attempt to learn Golang, Sugoku will eventually lead to both
text-based and graphical user interfaces.  In addition, re-immersing
myself in complexity and trying to optimize this NP-hard problem.  More
than the classic 9x9 grid will be supported!

-----------
Development
-----------

Roadmap
=======

#. Libraries written to be re-used by other projects
#. Text-based interface
#. GUI-based interface

-------
Objects
-------

* board/grid
  * cell
  * row
  * column
  * block/box
* character set
* ui rendering
  * GUI
  * TUI/CUI
* player
* solver
* checker
* hints

1x1:

* | * | *
--|---|--
* | * | *
--|---|--
* | * | *

2x2:

* * | * * | * *
* * | * * | * *
----|-----|----
* * | * * | * *
* * | * * | * *
----|-----|----
* * | * * | * *
* * | * * | * *

3x3:

* * * | * * * | * * *
* * * | * * * | * * *
* * * | * * * | * * *
------|-------|------
* * * | * * * | * * *
* * * | * * * | * * *
* * * | * * * | * * *
------|-------|------
* * * | * * * | * * *
* * * | * * * | * * *
* * * | * * * | * * *

4x4:

* * * * | * * * * | * * * *
* * * * | * * * * | * * * *
* * * * | * * * * | * * * *
* * * * | * * * * | * * * *
--------|---------|--------
* * * * | * * * * | * * * *
* * * * | * * * * | * * * *
* * * * | * * * * | * * * *
* * * * | * * * * | * * * *
--------|---------|--------
* * * * | * * * * | * * * *
* * * * | * * * * | * * * *
* * * * | * * * * | * * * *
* * * * | * * * * | * * * *

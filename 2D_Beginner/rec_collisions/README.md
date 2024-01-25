
# 2D - COLLISION RECTANGLES
Demonstrates how to check for collisions of moving rectangles against other moving rectangles as well as border rectangles. This method uses *rl.CheckCollisionRecs* and creates a duplicate rectangle which is then moved to the next position of the moving rectangle. If this next rectangle collides then the direction is changed. Checking for collisions with the **next** movement position, as opposed to the moving rectangle itself, prevents problems with rectangles intersecting. When using the moving rectangle itself, when the collision is reported as having happened, the rectangles are already intersecting which can cause problems. View more at [unklnik.com](https://unklnik.com/posts/2d-rectangle-collisions/)

https://github.com/unklnik/raylib-go-more-examples/assets/146096950/5773316b-27bd-4a98-b00d-f202be00dc3c

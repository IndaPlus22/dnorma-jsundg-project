# Project Specification
[Project](https://github.com/IndaPlus22/dnorma-jsundg-project)

## Platformer with Go
We aim to make a platformer with Golang. We'll make use of the [Pixel package](https://pkg.go.dev/github.com/faiface/pixel). The game will be in singleplayer and the playable character will be able to interact with objects and obstacles in different ways. The project is quite feasible since we use a package for the graphics which is the main concern. An MVP will be done by week 2 while the full completion of the game might not be as feasible.

We will be using a simpler version of Git Flow as our Git methodology. There will be a `main`branch which is protected from pushes - only merging pull requests can be done. Under `main` is the `develop` branch which merges all `feature` branches. This branch only merges with main once a major milestones has been reached. There will also be `feature` branches which handle one issue at the time. They all merge in to the `develop` branch when completed.

### The game
The game will have 2D graphics and the camera will move with the character. The character will be a computer science student in a cerise overall that gets powerups from potions(drinks). The win condition is to defeat all enemies and get to the exit. The player might be able to get achievements in the form of märken if there is time to implement it.
 
### Requirements
* There should be a playable character that can move around the map with the WASD keys
    * Player collisions has to be implemented
    * Getting rid of obstacles should be available in certain situations
* There should be obstacles that the player can interact with
    * "Enemies" with simple movement paths should exist
* The player should be able to win or lose
    * Scores and highscores should be kept
* Decent graphics on everything

## MVP
* Movable character
* Gravity
* Working collisions
* Win/lose condition
* Map design

## Team members
* Daniel Norman aka dnorma
* Johan Sundgren aka jsundg

## Naming conventions
##### Issues and commits
Issues and commit messages are going to be named according to github standards, in other words in future tense. Examples: "Add wall collision", "Fix annoying bug".

##### Pull requests
Pull requests will be named after their associated issue. It will be in the format of "issue/[issue number]-[issue-name], i.e issue #13 "Improve performance" would be "issue/13-improve-performance".

## Weekly milestones
Week 1:     Character movement
Week 2:     Walls and floors
Week 3:     Collisions & items  
Week 4:     Enemies & level design
Week 5:     Movable camera
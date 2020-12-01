let game;
let config;

config = {
  type: Phaser.AUTO,
  width: 800,
  height: 800,
  backgroundColor: '#3CB371',
  parent: 'game',
  scene: {
    preload: preload,
    create: create,
    update: update,
  }
}

let snake, cursors, gfood, rfood, score = 0, userID = '';
let scoreText, gameOverText, startAgain, userName = '';
const UP = 0, DOWN = 1, LEFT = 2, RIGHT = 3;

game = new Phaser.Game(config);

function preload () {
  this.load.image('gfood', './assets/gapple.png');
  this.load.image('body', './assets/body.png');
  this.load.image('rfood', './assets/rapple.png')
  this.load.image('obtacle', './assets/bom.png')
}

function create() {
  let gFood = new Phaser.Class({
    Extends: Phaser.GameObjects.Image,
    initialize:

    function gFood(scene, x, y) {
      Phaser.GameObjects.Image.call(this, scene);

      this.setTexture('gfood');
      this.setPosition(x * 16, y * 16);
      this.setOrigin(0);
      this.total = 0;

      scene.children.add(this);
    },

    geat: function() {
      this.total++;
    }
  });

  let rFood = new Phaser.Class({
    Extends: Phaser.GameObjects.Image,
    initialize:

    function rFood(scene, x, y) {
      Phaser.GameObjects.Image.call(this, scene);

      this.setTexture('rfood');
      this.setPosition(x * 16, y * 16);
      this.setOrigin(0);
      this.total = 0;

      scene.children.add(this);
    },

    reat: function() {
      this.total--;
    },
  });

  let Snake = new Phaser.Class({
    initialize:

    function Snake(scene, x, y) {
      this.headPosition = new Phaser.Geom.Point(x, y);
      this.body = scene.add.group();

      this.head = this.body.create(x * 16, y * 16, 'body');
      this.head.setOrigin(0);

      this.alive = true;

      this.speed = 100;

      this.moveTime = 0;

      this.tail = new Phaser.Geom.Point(x, y);

      this.heading = RIGHT;
      this.direction = RIGHT;
    },

    update: function(time) {
      if (time >= this.moveTime) {
        return this.move(time);
      }
    },

    Left: function() {
      if (this.direction === UP || this.direction === DOWN) {
        this.heading = LEFT;
      }
    },

    Right: function() {
      if (this.direction === UP || this.direction === DOWN) {
        this.heading = RIGHT;
      }
    },

    Up: function() {
      if (this.direction === LEFT || this.direction === RIGHT) {
        this.heading = UP;
      }
    },

    Down: function() {
      if (this.direction === LEFT || this.direction === RIGHT) {
        this.heading = DOWN;
      }
    },

    move: function(time) {
      switch (this.heading) {
        case LEFT:
          this.headPosition.x = Phaser.Math.Wrap(this.headPosition.x - 1, 0, 50);
          break;
        case RIGHT:
          this.headPosition.x = Phaser.Math.Wrap(this.headPosition.x + 1, 0, 50);
          break;
        case UP:
          this.headPosition.y = Phaser.Math.Wrap(this.headPosition.y - 1, 0, 50);
          break;
        case DOWN:
          this.headPosition.y = Phaser.Math.Wrap(this.headPosition.y + 1, 0, 50);
          break;
      }

      this.direction = this.heading;
      // Updating the body segment
      Phaser.Actions.ShiftPosition(this.body.getChildren(), this.headPosition.x * 16, this.headPosition.y * 16, 1, this.tail);

      let hitBody = Phaser.Actions.GetFirst(this.body.getChildren(), { x: this.head.x, y: this.head.y }, 1);

      if (hitBody) {
        this.alive = false;
        gameOverText.visible = true;

        this.updateScore()

        return false;
      } else {
        // Updating the timer ready for the next movement
        this.moveTime = time + this.speed;
      }
      return true;
    },

    updateScore: function() {
      fetch('http://localhost:8000/users', {mode: 'cors'})
      .then(function(response) {
        return response.json();
      })
      .then(function(text) {
        userName = text[0].username;
        userID = text[0].user_id
        console.log('Request successful', userName, userID);

        fetch('http://localhost:8000/users/' + userID, {
          method: 'PUT',
          headers: { 'content-type': 'application/json' },
          body: JSON.stringify({
            score: score.toString()
          })
        });
      })
      .catch(function(error) {
        log('Request failed', error);
      });
    },

    grow: function() {
      let large = this.body.create(this.tail.x, this.tail.y, 'body');

      large.setOrigin(0);
    },

    getGreenFood: function(gfood) {
      if (this.head.x === gfood.x && this.head.y === gfood.y) {
        this.grow();
        gfood.geat();

        // For every multiple of 5 the speed is aumented
        if (this.speed > 20 && gfood.total % 5 === 0) {
          this.speed -= 10;
        }

        score += 5;
        scoreText.setText('Score: ' + score);
        return true;
      }
      else {
        return false;
      }
    },

    getRedFood: function(rfood) {
      if (this.head.x === rfood.x && this.head.y === rfood.y) {
        rfood.reat();

        // For every multiple of 5 the speed is aumented
        if (this.speed > 20 && rfood.total % 5 === 0) {
          this.speed -= 15;
        }

        score -= 10;
        scoreText.setText('Score: ' + score);

        if (score <= 0) {
          this.alive = false;
          gameOverText.visible = true;
  
          this.updateScore()
        }
        return true;
      }
      else {
        return false;
      }
    },

    updateGrid: function(grid) {
      this.body.children.each(function (segment) {
        let bodyX = segment.x / 16;
        let bodyY = segment.x / 16;

        grid[bodyY][bodyX] = false;
      });

      return grid;
    },
  });

  snake = new Snake(this, 25, 25);
  gfood = new gFood(this, 3, 4);
  rfood = new rFood(this, 12, 16);

  // Score
  textStyleScore = { font: "bold 20px sans-serif", fill: "white", align: "center" };
  textGameOver = { font: "bold 64px sans-serif", fill: "white", align: "center" };
  scoreText = this.add.text(30, 20, "Score: 0", textStyleScore);
  gameOverText = this.add.text(200, 350, "GAME OVER", textGameOver);
  gameOverText.setOrigin(0);
  gameOverText.visible = false;

  // Keyboard controls
  cursors = this.input.keyboard.createCursorKeys();
}

function update (time, delta) {
  if (!snake.alive) {
    return;
  }

  if (cursors.left.isDown) {
      snake.Left();
  }
  else if (cursors.right.isDown) {
      snake.Right();
  }
  else if (cursors.up.isDown) {
      snake.Up();
  }
  else if (cursors.down.isDown) {
      snake.Down();
  }

  if (snake.update(time)) {
    if (snake.getGreenFood(gfood)) {
      foodReposition(gfood);
    }
    else if (snake.getRedFood(rfood)) {
      foodReposition(rfood);
    }
  }
}

function foodReposition(food) {
  let positions = []

  for (let y = 0; y < 50; y++) {
    positions[y] = [];

    for (let x = 0; x < 50; x++) {
      positions[x] = true;
    }
  }

  snake.updateGrid(positions);

  let validLocations = []
  // Remove false positions
  for (let y = 0; y < 50; y++) {
    positions[y] = [];

    for (let x = 0; x < 50; x++) {
      if (positions[x] === true) {
        // If position is valid, add it.
        validLocations.push({x: x, y: y});
      }
    }
  }

  if (validLocations.length > 0) {
    // Pick a random food position
    let pos = Phaser.Math.RND.pick(validLocations);
    // Place it the food
    food.setPosition(pos.x * 16, pos.y * 16);

    return true;
  } else {
    return false;
  }
}

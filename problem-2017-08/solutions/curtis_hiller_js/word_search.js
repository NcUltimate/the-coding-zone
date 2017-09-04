var reader = require('fs');

(function () {
  var size;
  var lines = [];
  var words = [];
  var highlights = []
  var right = left = up = down = upleft = downleft = upright = downright = '';

  function output(){
    var rows = highlights.length;
    var cols;
    var i = 0;
    var j;
    var output;

    for (; i < rows; i++){
      output = '';
      j = 0;
      for (; j < highlights[i].length; j++){
        output += highlights[i][j];
      }
      console.log(output);
    }
  }

  function highlightLetter(y, x){
    newLetter  = '\x1b[44m\x1b[30m';
    newLetter += lines[y].charAt(x);
    newLetter += '\x1b[0m\x1b[0m';
    highlights[y][x] = newLetter;
  }

  function findWordLetters (direction, idx, word){
    var length = word.length;
    var x, y;
    var delta_x = delta_y = 0;

    switch (direction){
      case 'right':
        delta_x = 1;
        x = idx % size[1];
        y = Math.floor(idx / size[1]);
      break;

      case 'left':
        delta_x = -1;
        y = Math.floor(idx / size[1]);
        x = Math.abs(size[0] - (idx % size[1]));
      break;

      case 'up':
        delta_y = -1;
        x = Math.floor(idx / size[0]);
        y = idx % size[1];
      break;

      case 'down':
        delta_y = 1;
        x = idx % size[0];
        y = Math.floor(idx / size[0]);
      break;

      default:
        return false;
      break;
    }


    // console.log('highlighting', direction, idx, word);
    // console.log('x', x, 'y', y);

    var i = 0;
    for(; i < length; i++){
      highlightLetter(y + (delta_y * i), x + (delta_x * i));
      // console.log(y + (delta_y * i), x + (delta_x * i));
      // console.log(highlights[y + (delta_y * i)]);
    }
  }

  function findWord(word){
    var idx;

    idx = up.indexOf(word);
    if (idx !== -1){
      return findWordLetters('up', idx, word);
    }

    idx = down.indexOf(word);
    if (idx !== -1){
      return findWordLetters('down', idx, word);
    }

    idx = left.indexOf(word);
    if (idx !== -1){
      return findWordLetters('left', idx, word);
    }

    idx = right.indexOf(word);
    if (idx !== -1){
      return findWordLetters('right', idx, word);
    }

    return false;
  }

  function solve (){
    var run = words.length;
    var i = 0;
    for (; i < run; i++){
      findWord(words[i]);
    }
  }

  function handleFileLoad (err, data){
    lines = data.split('\n');

    size = lines.shift().split(' ');
    size[0] = parseInt(size[0]);
    size[1] = parseInt(size[1]);

    while (lines.length > size[0]){
      words.push(lines.pop());
    }

    var i = 0;
    for (; i < lines.length; i++){
      highlights[i] = lines[i].split('');
    }

    right = lines.join('');
    left = right.split('').reverse().join('');

    var i = 0;
    var j, k;
    for (; i < size[1]; i++){
      j = 0;
      k = size[0] - 1;
      for (; j < size[0]; j++, k--){
        down += lines[j].charAt(i);
        up += lines[k].charAt(i);
      }
    }

    solve();
    output();
  }

  function initialize (fileName){
    reader.onload = handleFileLoad;
    reader.readFile(fileName, 'utf8', handleFileLoad);
  }

  initialize('input.txt');
}());

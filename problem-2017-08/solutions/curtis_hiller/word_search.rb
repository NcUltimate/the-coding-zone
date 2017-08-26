class WordSearch

  attr_accessor :size, :letters, :rows, :words, :highlights
  def initialize(file_path = '')
    @letters = {}
    @words = []
    @highlights = []
    ingest_file!(file_path)
    solve
    output_colorized
  end

  private

  def ingest_file!(file_path)
    begin
      f = File.open(file_path, 'r')
    rescue
      raise 'Could not open file.'
    end

    temp =  f.readline.split(' ').map(&:to_i)
    @size = [temp[1], temp[0]]
    @rows = Array.new(size[0])

    y = 0
    f.each_line do |line|
      line.split('').each_with_index do |letter, x|
        rows[x] ||= []
        rows[x][y] = letter
        letters[letter] ||= []
        letters[letter] << [x, y]
      end
      y += 1
      break if y == size[1]
    end

    f.each_line { |line| words << line.to_s.gsub(/\s/,'') }
  end

  def find_letter(x, y, letter)
    return unless x >= 0 && y >= 0
    return unless x < size[0] && y < size[1]
    return true if letter == rows[x][y]
  end

  def path(xy, delta_x, delta_y, word_letters)
    x = xy[0]
    y = xy[1]
    idx = 1
    path = []
    path << [x, y]
    x += delta_x
    y += delta_y
    while find_letter(x, y, word_letters[idx])
      path << [x, y]
      return path if path.length == word_letters.length
      x += delta_x
      y += delta_y
      idx += 1
    end
    false
  end

  def go(xy, word_letters)
    # up left (-1, 1)
    new_path = path(xy, -1, 1, word_letters)
    return new_path if new_path

    # down (0, 1)
    new_path = path(xy, 0, 1, word_letters)
    return new_path if new_path

    # down right (1, 1)
    new_path = path(xy, 1, 1, word_letters)
    return new_path if new_path

    # right (1, 0)
    new_path = path(xy, 1, 0, word_letters)
    return new_path if new_path

    # up right (1, -1)
    new_path = path(xy, 1, -1, word_letters)
    return new_path if new_path

    # up (0, -1)
    new_path = path(xy, 0, -1, word_letters)
    return new_path if new_path

    # up left (-1, -1)
    new_path = path(xy, -1, -1, word_letters)
    return new_path if new_path

    # left (-1, 0)
    new_path = path(xy, -1, 0, word_letters)
    return new_path if new_path
  end

  def solve
    words.each do |word|
      word_letters = word.split('')
      letters[word_letters[0]].each do |xy|
        results = go(xy, word_letters)
        results.each { |res| highlights << res } if results
      end
    end
  end

  def output_colorized
    y = 0
    while y < size[1]
      x = 0
      output = ''
      while x < size[0]
        letter = rows[x][y]
        output <<
          if highlights.include?([x, y])
            "\e[33m#{letter}\e[0m"
          else
            letter
          end
        x += 1
      end
      puts output
      y += 1
    end
  end
end

WordSearch.new('input.txt')
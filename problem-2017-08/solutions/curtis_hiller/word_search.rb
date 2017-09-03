class Descartes
  attr_accessor :x, :y
  def initialize(x, y)
    @x = x
    @y = y
  end

  def coordinates
    [x, y]
  end
end

class WordSearch
  attr_accessor :size, :letters, :rows, :words, :highlights
  def initialize(file_path = '')
    @letters = {}
    @words = []

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

    size_array =  f.readline.split(' ').map(&:to_i)
    @size = [size_array[1], size_array[0]]
    @rows = Array.new(size[0])
    @highlights = Array.new(size[0])

    y = 0
    f.each_line do |line|
      line = line.gsub(/[^A-Z]/, '')
      line.split('').each_with_index do |letter, x|
        rows[x] ||= []
        rows[x][y] = letter
        highlights[x] ||= []
        highlights[x][y] = letter
        letters[letter] ||= []
        letters[letter] << Descartes.new(x, y)
      end
      y += 1
      break if y == size[1]
    end

    f.each_line { |line| words << line.to_s.gsub(/[^A-Z]/, '') }
  end

  def location_has_letter?(coordinates, letter)
    return unless coordinates.x >= 0 && coordinates.y >= 0
    return unless coordinates.x < size[0] && coordinates.y < size[1]
    return true if letter == rows[coordinates.x][coordinates.y]
  end

  def follow_word_path(coordinates, delta_x, delta_y, word_letters)
    path = [coordinates]

    next_coordinates =
      Descartes.new(coordinates.x + delta_x, coordinates.y + delta_y)

    idx = 1
    while location_has_letter?(next_coordinates, word_letters[idx])
      path << next_coordinates
      return path if path.length == word_letters.length
      next_coordinates = Descartes.new(
        next_coordinates.x + delta_x,
        next_coordinates.y + delta_y
      )
      idx += 1
    end

    false
  end

  def find_word_path(coordinates, word_letters)
    word_length = word_letters.length

    # left and...
    if coordinates.x >= word_length
      # only left (-1, 0)
      new_path = follow_word_path(coordinates, -1, 0, word_letters)
      return new_path if new_path

      # down (-1, 1)
      if coordinates.y <= size[1] - word_length
        new_path = follow_word_path(coordinates, -1, 1, word_letters)
        return new_path if new_path
      end

      # up (-1, -1)
      if coordinates.y >= word_length - 1
        new_path = follow_word_path(coordinates, -1, -1, word_letters)
        return new_path if new_path
      end
    end

    # right and...
    if coordinates.x <= size[0] - word_length
      # only right
      new_path = follow_word_path(coordinates, 1, 0, word_letters)
      return new_path if new_path

      # down (1, 1)
      if coordinates.y <= size[1] - word_length
        new_path = follow_word_path(coordinates, 1, 1, word_letters)
        return new_path if new_path
      end

      # up (1, -1)
      if coordinates.y >= word_length - 1
        new_path = follow_word_path(coordinates, 1, -1, word_letters)
        return new_path if new_path
      end
    end

    # down (0, 1)
    if coordinates.y <= size[1] - word_length
      new_path = follow_word_path(coordinates, 0, 1, word_letters)
      return new_path if new_path
    end

    # up (0, -1)
    if coordinates.y >= word_length - 1
      new_path = follow_word_path(coordinates, 0, -1, word_letters)
      return new_path if new_path
    end

    []
  end

  def highlight_letters(letters)
    letters.each do |coordinates|
      highlights[coordinates.x][coordinates.y] =
        "\e[30m\e[46m#{rows[coordinates.x][coordinates.y]}\e[0m\e[0m"
    end
  end

  def solve
    words.each do |word|
      word_letters = word.split('')
      letters[word_letters[0]].each do |coordinates|
        highlight_letters(find_word_path(coordinates, word_letters))
      end
    end
  end

  def output_colorized
    y = 0
    while y < size[1]
      x = 0
      output = ''
      while x < size[0]
        output << highlights[x][y]
        x += 1
      end
      puts output
      y += 1
    end
  end
end

WordSearch.new('input.txt')

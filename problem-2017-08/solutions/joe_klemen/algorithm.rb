module WordSearch
  class Algorithm
    def self.run!(grid, words)
      self.new(grid, words).run!
    end

    attr_accessor :grid, :words, :locations, :added
    def initialize(grid, words)
      self.grid = grid.map(&:upcase).map(&:chars)
      self.words = words.map(&:upcase)
      self.locations = {}
      self.added = 0
    end

    def run!
      words = self.words.sort_by(&:length).reverse!

      marked = {}
      words.each do |word|
        idx = index_of(word, marked)
        next unless idx

        word.length.times do |k|
          key = to_key(idx + k)
          next if marked[key]

          marked[key] = true
          coord   = locations[key]
          letter  = grid[coord.x][coord.y]
          grid[coord.x][coord.y] = red(letter)
        end
      end
      grid.map(&:join).join("\n")
    end

    private

    def index_of(word, marked)
      overlapping, idx = true, -1
      while overlapping
        start = idx
        idx = search_str.index(word, start + 1)
        break unless idx

        overlapping = (idx...(idx + word.length)).all? do |k|
          marked[to_key(k)]
        end
      end
      idx
    end

    def to_key(k)
      orig_len = (search_str.length / 2)
      return k if k < orig_len

      orig_len - ((k) % orig_len)
    end

    def red(char)
      "\e[31m#{char}\e[0m"
    end

    def horizontal_search_str
      horizontal = ''
      (0...grid.length).each do |row|
        (0...grid[0].length).each do |col|
          horizontal << grid[row][col]
          locations[self.added] = WordSearch::Point[row, col]
          self.added += 1
        end
        horizontal << ' '
        self.added += 1
      end
      horizontal.chomp(' ')
    end

    def vertical_search_str
      vertical = ''
      (0...grid[0].length).each do |col|
        (0...grid.length).each do |row|
          vertical << grid[row][col]
          locations[self.added] = WordSearch::Point[row, col]
          self.added += 1
        end
        vertical << ' '
        self.added += 1
      end
      vertical.chomp(' ')
    end

    def diagonal_search_str
      diagonal = ''
      sr, sc = 0, 0
      until sr == grid.length
        r, c = sr, sc
        until r == grid.length || c < 0
          diagonal << grid[r][c]
          locations[self.added] = WordSearch::Point[r, c]
          self.added += 1
          r += 1
          c -= 1
        end

        if sc == grid[0].length - 1
          sr += 1
          diagonal << ' '
          self.added += 1
        end

        if sc < grid[0].length - 1
          sc += 1 
          diagonal << ' '
          self.added += 1
        end
      end
      diagonal.chomp(' ')
    end

    def antidiagonal_search_str
      antidiagonal = ''
      sr, sc = 0, grid[0].length - 1
      until sc < 0
        r, c = sr, sc
        until r < 0 || c < 0
          antidiagonal << grid[r][c]
          locations[self.added] = WordSearch::Point[r, c]
          self.added += 1
          r -= 1
          c -= 1
        end

        if sr == grid.length - 1
          sc -= 1
          antidiagonal << ' '
          self.added += 1
        end

        if sr < grid.length - 1
          sr += 1 
          antidiagonal << ' '
          self.added += 1
        end
      end
      antidiagonal.chomp(' ')
    end

    def search_str
      return @search_str if @search_str
      # add horizontal
      @search_str = horizontal_search_str
      @search_str << ' '

      # add vertical
      @search_str << vertical_search_str
      @search_str << ' '

      # add diagonal
      @search_str << diagonal_search_str
      @search_str << ' '

      # add antidiagonal
      @search_str << antidiagonal_search_str

      # add reversal of everything
      temp = @search_str.reverse
      @search_str << ' ' << temp
    end
  end
end
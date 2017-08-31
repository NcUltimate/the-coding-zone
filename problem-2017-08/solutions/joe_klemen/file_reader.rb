module WordSearch  
  class FileReader
    class << self
      def read!(filename)
        self.new(filename).read!
      end
    end

    attr_accessor :filename
    def initialize(filename)
      @filename = filename
    end

    def read!
      [
        grid,
        words
      ]
    end

    private

    def to_grid(coord)
      coord[0] = coord[0].to_i
      coord[1] = coord[1].to_i
      coord
    end

    def lines
      @lines ||= File.open(filename, 'r') do |f|
        f.readlines.map(&:chomp)
      end
    end

    def grid
      lines[1..dimensions[0]]
    end

    def words
      lines[(dimensions[0] + 1)..-1]
    end

    def dimensions
      @dimensions ||= lines[0].split(/s/).map(&:to_i)
    end
  end
end
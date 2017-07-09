module Mario  
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
        Mario::Point[*dimensions],
        Mario::Point[*player_start],
        blocks,
        coins
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
        f.readlines.map do |line|
          line.chomp.split(/\s+/)
        end
      end
    end

    def dimensions
      @dimensions ||= lines[0].map(&:to_i)
    end

    def player_start
      @player_start ||= to_grid(lines[1])
    end

    def blocks
      @blocks ||= locations_for('B')
    end

    def coins
      @coins ||= locations_for('C')
    end

    def locations_for(type)
      return {} if locations[type].nil?

      locations[type].each_with_object({}) do |loc, hash|
        x, y, mark = to_grid(loc)
        hash[x] ||= {}
        hash[x][y] = mark
      end
    end

    def locations
      @locations ||= lines[2..-1].group_by { |loc| loc[2] }
    end
  end
end
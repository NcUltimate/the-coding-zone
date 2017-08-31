module WordSearch
  class Point
    class << self
      def [](x, y)
        self.new(x, y)
      end
    end

    attr_accessor :x, :y
    def initialize(x, y)
      self.x = x
      self.y = y
    end

    def inspect
      "(#{x}, #{y})"
    end

    def to_s
      inspect
    end
  end
end
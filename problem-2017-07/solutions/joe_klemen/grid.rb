module Mario
  class Grid
    attr_accessor :dimensions, :start, :blocks, :coins
    def initialize(dimensions, start, blocks = {}, coins = {})
      self.dimensions = dimensions
      self.start = start
      self.blocks = blocks
      self.coins = coins
    end

    def walkable_spaces
      @wspaces ||= walkable_block_spaces + walkable_bottom_spaces
    end

    def block_at?(p)
      blocks[p.x] && blocks[p.x][p.y]
    end

    def coin_at?(p)
      coins[p.x] && coins[p.x][p.y]
    end

    def include?(point)
      point.x > -1 && point.x < width &&
        point.y > -1 && point.y < height
    end

    def can_enter?(p)
      include?(p) && !block_at?(p)
    end

    def width
      dimensions.x
    end

    def height
      dimensions.y
    end

    def size
      width * height
    end

    def inspect
      "#<Grid #{height}x#{width}>"
    end

    private

    def walkable_block_spaces
      @bl_spaces ||= blocks.reduce([]) do |ary, (x, yh)|
        yh.keys.each do |y|
          next ary unless include?(Point[x, y + 1])
          next ary if block_at?(Point[x, y + 1])

          ary << Point[x, y + 1]
        end
        ary
      end
    end

    def walkable_bottom_spaces
      @bo_spaces ||= width.times.reduce([]) do |ary, x|
        next ary if block_at?(Point[x, 0])
        ary << Point[x, 0]
      end
    end
  end
end
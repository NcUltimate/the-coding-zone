module Intersections
  class Line < Shape
    attr_accessor :p1, :p2
    def initialize(p1, p2)
      self.p1 = p1
      self.p2 = p2
    end

    def slope
      @slope ||= (p2.y - p1.y).to_f / (p2.x - p1.x).to_f
    end

    def offset_y
      @offset_y ||= p2.y - slope * p2.x
    end

    def midpoint
      @midpoint ||= Point.new(
        (p1.x + p2.x).to_f / 2.0,
        (p1.y + p2.y).to_f / 2.0
      )
    end

    def min_x
      p1.x <= p2.x ? p1.x : p2.x
    end

    def min_y
      p1.y <= p2.y ? p1.y : p2.y
    end

    def max_x
      p1.x >= p2.x ? p1.x : p2.x
    end

    def max_y
      p1.y >= p2.y ? p1.y : p2.y
    end

    def f(x)
      slope * x + offset_y
    end

    def intersection(line)
      return if slope == line.slope && slope.abs == Float::INFINITY
      int = 
        if line.slope.abs == Float::INFINITY
          x = line.p1.x
          Point.new(x, f(x))
        elsif slope.abs == Float::INFINITY
          x = p1.x
          Point.new(x, line.f(x))
        else
          x = (line.offset_y - offset_y).to_f / (slope - line.slope).to_f
          Point.new(x, f(x))
        end

      return int if bounds_point?(int) && line.bounds_point?(int)
    rescue
      nil
    end

    def split_by?(point)
      intersects_point?(point) &&
        !point.eql?(p1) &&
        !point.eql?(p2)
    end

    def bounds_point?(point)
      point.x >= min_x &&
        point.x <= max_x &&
        point.y >= min_y &&
        point.y <= max_y
    end

    def intersects_point?(point)
      return false unless bounds_point?(point)
      if slope.abs == Float::INFINITY
        point.x == p1.x
      else
        f(point.x) == point.y
      end
    end

    def intersects_line?(line)
      !intersection(line).nil?
    end

    def intersects_poly?(poly)
      poly.intersects_line?(self)
    end

    def to_s
      "#<Line (#{p1.x}, #{p1.y}) --> (#{p2.x}, #{p2.y})>"
    end

    def inspect
      to_s
    end
  end
end
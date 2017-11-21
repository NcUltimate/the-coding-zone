module Intersections
  class Shape
    def intersects?(shape)
      if shape.is_a?(Point)
        intersects_point?(shape)
      elsif shape.is_a?(Line)
        intersects_line?(shape)
      elsif shape.is_a?(Polygon)
        intersects_poly?(shape)
      end
    end

    def intersects_point?(point)
      false
    end

    def intersects_line?(line)
      false
    end

    def intersects_poly?(poly)
      false
    end
  end
end
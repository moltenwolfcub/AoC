package Helpers;

public class Position2D implements Cloneable {
    public Integer x;
    public Integer y;

    public Position2D() {
        this.x = 0;
        this.y = 0;
    }

    public Position2D(Integer x, Integer y) {
        this.x = x;
        this.y = y;
    }

    @Override
    public Position2D clone() {
        return new Position2D(x, y);
    }

    @Override
    public String toString() {
        return "<x: "+ this.x+", y: "+this.y+">";
    }

    @Override
    public int hashCode() {
        final int prime = 31;
        int hash = 7;
        hash = prime * hash + ((x == null) ? 0 : x.hashCode());
        hash = prime * hash + ((y == null) ? 0 : y.hashCode());
        return hash;
    }

    @Override
    public boolean equals(Object other) {
        if (this == other) {
            return true;
        }
        if (other == null || this == null) {
            return false;
        }
        if (other instanceof Position2D) {
            Position2D pos = (Position2D) other;
            return pos.x == this.x && pos.y == this.y;
        }
        return false;
    }

}

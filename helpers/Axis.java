package helpers;

public enum Axis {
    X,
    Y;

    private Axis() {

    }

    public Position2D movePosition(Position2D originalPos, Integer delta) {
        if(this.equals(Axis.X)) {
            return new Position2D(originalPos.x+delta, originalPos.y);
        } else if(this.equals(Axis.Y)) {
            return new Position2D(originalPos.x, originalPos.y+delta);
        } else {
            throw new IllegalArgumentException("Can't move a 2d position by any axis except X or Y.");
        }
    }

    public Integer getPositionAxisValue(Position2D pos) {
        if(this.equals(Axis.X)) {
            return pos.x;
        } else if(this.equals(Axis.Y)) {
            return pos.y;
        } else {
            throw new IllegalArgumentException("Can't get any axis except X or Y from a 2d position.");
        }
    }

    public Axis getOtherAxis2D() {
        if(this.equals(Axis.X)) {
            return Axis.Y;
        } else if(this.equals(Axis.Y)) {
            return Axis.X;
        } else {
            throw new IllegalArgumentException("Current Axis isn't a 2D Axis.");
        }
    }
}

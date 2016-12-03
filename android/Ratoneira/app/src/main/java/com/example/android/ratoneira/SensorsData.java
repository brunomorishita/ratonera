package com.example.android.ratoneira;

import java.util.Date;

public class SensorsData {

    private Date timestamp;
    private Accelerometer acc;

    public SensorsData(Date timestamp, float acc_x, float acc_y, float acc_z) {
        this.timestamp = timestamp;

        acc = new Accelerometer(acc_x, acc_y, acc_z);
       /* this.acc.setX(acc_x);
        this.acc.setY(acc_y);
        this.acc.setZ(acc_z);*/
    }

    public Accelerometer getAcc() {
        return acc;
    }

    public Date getTimestamp() {
        return timestamp;
    }

    // Accelerometer Data
    public class Accelerometer {
        private float x;
        private float y;
        private float z;
        // get and set

        public float getX() {
            return x;
        }

        public float getZ() {
            return z;
        }

        public float getY() {
            return y;
        }

/*        public void setX(float x) {
            this.x = x;
        }

        public void setY(float y) {
            this.y = y;
        }

        public void setZ(float z) {
            this.z = z;
        }*/

        Accelerometer(float x, float y, float z){
            this.x = x;
            this.y = y;
            this.z = z;
        }
    }
}
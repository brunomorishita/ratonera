package com.example.android.ratoneira;


public class SensorsData {

    private long timestamp;
    private Accelerometer acc;
    private GPS gps;
    private String deviceId;

    public SensorsData(String deviceId, long timestamp, float acc_x, float acc_y, float acc_z, int lat, int lgt) {
        this.timestamp = timestamp;
        this.deviceId = deviceId;

        acc = new Accelerometer(acc_x, acc_y, acc_z);
       /* this.acc.setX(acc_x);
        this.acc.setY(acc_y);
        this.acc.setZ(acc_z);*/
        gps = new GPS(lat, lgt);
    }

    public Accelerometer getAcc() {
        return acc;
    }

    public GPS getGPS() {
        return gps;
    }

    public long getTimestamp() {
        return timestamp;
    }

    public String getDeviceId() {return deviceId; }

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

    public class GPS {
        private int lat;
        private int lgt;
        // get and set

        public float getLat() {
            return lat;
        }

        public float getLgt() {
            return lgt;
        }

        GPS(int lat, int lgt){
            this.lat = lat;
            this.lgt = lgt;
        }
    }
}
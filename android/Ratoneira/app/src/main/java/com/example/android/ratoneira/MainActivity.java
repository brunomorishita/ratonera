package com.example.android.ratoneira;

import android.Manifest;
import android.content.Context;
import android.content.pm.PackageManager;
import android.hardware.Sensor;
import android.hardware.SensorEvent;
import android.hardware.SensorEventListener;
import android.hardware.SensorManager;
import android.location.Criteria;
import android.location.Location;
import android.location.LocationListener;
import android.location.LocationManager;
import android.os.Build;
import android.os.Handler;
import android.support.v4.app.ActivityCompat;
import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;
import android.telephony.TelephonyManager;
import android.util.Log;
import android.widget.TextView;
import android.widget.EditText;
import android.view.View;
import android.widget.Toast;

import java.net.URI;
import java.net.URISyntaxException;
import java.util.Date;

import org.java_websocket.client.WebSocketClient;
import org.java_websocket.drafts.Draft_17;
import org.java_websocket.handshake.ServerHandshake;

public class MainActivity extends AppCompatActivity implements SensorEventListener, LocationListener {

    private Handler handler = new Handler();

    private float acc_x, acc_y, acc_z;

    private int lat, lng;

    private WebSocketClient mWebSocketClient;

    private LocationManager locationManager;
    private String provider;
    private SensorManager senSensorManager;
    private Sensor senAccelerometer;
    String deviceId;

    private String cachedServerURI=null;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);

        TelephonyManager telephonyManager = (TelephonyManager)getSystemService(Context.TELEPHONY_SERVICE);
        deviceId = telephonyManager.getDeviceId();

        // Initialize acceleromenter
        senSensorManager = (SensorManager) getSystemService(Context.SENSOR_SERVICE);
        senAccelerometer = senSensorManager.getDefaultSensor(Sensor.TYPE_ACCELEROMETER);
        senSensorManager.registerListener(this, senAccelerometer, SensorManager.SENSOR_DELAY_NORMAL);

        // Get the location manager
        locationManager = (LocationManager) getSystemService(Context.LOCATION_SERVICE);
        // Define the criteria how to select the locatioin provider -> use
        // default
        Criteria criteria = new Criteria();
        provider = locationManager.getBestProvider(criteria, false);

        handler.postDelayed(runnable, 1000);
    }

    private void connectWebSocket(String serverURI) {
        URI uri;
        try {
            uri = new URI(serverURI);
            cachedServerURI = serverURI;
        } catch (URISyntaxException e) {
            e.printStackTrace();
            return;
        }

        mWebSocketClient = new WebSocketClient(uri, new Draft_17()) {
            @Override
            public void onOpen(ServerHandshake serverHandshake) {
                Log.i("Websocket", "Opened");
                mWebSocketClient.send("Hello from " + Build.MANUFACTURER + " " + Build.MODEL);
            }

            @Override
            public void onMessage(String s) {
                final String message = s;
                runOnUiThread(new Runnable() {
                    @Override
                    public void run() {
                        TextView textView = (TextView) findViewById(R.id.messages);
                        textView.setText(textView.getText() + "\n" + message);
                    }
                });
            }

            @Override
            public void onClose(int i, String s, boolean b) {
                Log.i("Websocket", "Closed " + s);
            }

            @Override
            public void onError(Exception e) {
                Log.i("Websocket", "Error " + e.getMessage());
            }
        };
        mWebSocketClient.connect();
    }

    protected void onPause() {
        super.onPause();
        senSensorManager.unregisterListener(this);

        // Remove thread
        handler.removeCallbacks(runnable);

        if (ActivityCompat.checkSelfPermission(this, Manifest.permission.ACCESS_FINE_LOCATION) != PackageManager.PERMISSION_GRANTED && ActivityCompat.checkSelfPermission(this, Manifest.permission.ACCESS_COARSE_LOCATION) != PackageManager.PERMISSION_GRANTED) {
            // TODO: Consider calling
            //    ActivityCompat#requestPermissions
            // here to request the missing permissions, and then overriding
            //   public void onRequestPermissionsResult(int requestCode, String[] permissions,
            //                                          int[] grantResults)
            // to handle the case where the user grants the permission. See the documentation
            // for ActivityCompat#requestPermissions for more details.
            return;
        }
        //locationManager.removeUpdates(this);

        if (mWebSocketClient != null && !mWebSocketClient.getConnection().isClosed()) {
            mWebSocketClient.close();
        }
    }

    protected void onResume() {
        super.onResume();
        if (ActivityCompat.checkSelfPermission(this, Manifest.permission.ACCESS_FINE_LOCATION) != PackageManager.PERMISSION_GRANTED && ActivityCompat.checkSelfPermission(this, Manifest.permission.ACCESS_COARSE_LOCATION) != PackageManager.PERMISSION_GRANTED) {
            // TODO: Consider calling
            //    ActivityCompat#requestPermissions
            // here to request the missing permissions, and then overriding
            //   public void onRequestPermissionsResult(int requestCode, String[] permissions,
            //                                          int[] grantResults)
            // to handle the case where the user grants the permission. See the documentation
            // for ActivityCompat#requestPermissions for more details.
            return;
        }
        locationManager.requestLocationUpdates(provider, 400, 1, this);
        senSensorManager.registerListener(this, senAccelerometer, SensorManager.SENSOR_DELAY_NORMAL);
        if (mWebSocketClient != null && mWebSocketClient.getConnection().isClosed()) {
            if (!cachedServerURI.isEmpty()) {
                connectWebSocket(cachedServerURI);
            }
        }
        handler.postDelayed(runnable,1000);
    }

    private boolean messageWasShown = false;
    private void sendWebsocketMessage(String message) {
        if (mWebSocketClient != null && mWebSocketClient.getConnection().isOpen()) {
            mWebSocketClient.send(message);
            messageWasShown = false;
        } else {
            if (!messageWasShown)
            Toast.makeText(MainActivity.this,
                    "Websocket connection is closed", Toast.LENGTH_SHORT).show();
        }
    }

    public void connect(View view) {
        EditText editText = (EditText) findViewById(R.id.message);
        connectWebSocket(editText.getText().toString());
    }


    private Runnable runnable = new Runnable() {
        @Override
        public void run() {
      /* do what you need to do */
            long time = new Date().getTime();
            SensorsData d = new SensorsData(deviceId, time, acc_x, acc_y, acc_z, lat, lng);
            String json = JsonUtil.toJSon(d);
            sendWebsocketMessage(json);
      /* and here comes the "trick" */
            handler.postDelayed(this, 1000);
        }
    };


    @Override
    public void onSensorChanged(SensorEvent sensorEvent) {
        Sensor mySensor = sensorEvent.sensor;

        // double check if it is accelerometer
        if (mySensor.getType() == Sensor.TYPE_ACCELEROMETER) {
            acc_x = sensorEvent.values[0];
            acc_y = sensorEvent.values[1];
            acc_z = sensorEvent.values[2];
        }
    }

    @Override
    public void onAccuracyChanged(Sensor sensor, int i) {

    }

    private boolean isLocationEnabled() {
        return locationManager.isProviderEnabled(LocationManager.GPS_PROVIDER) ||
                locationManager.isProviderEnabled(LocationManager.NETWORK_PROVIDER);
    }

    @Override
    public void onLocationChanged(Location location) {
        lat = (int) (location.getLatitude());
        lng = (int) (location.getLongitude());
    }

    @Override
    public void onStatusChanged(String s, int i, Bundle bundle) {

    }

    @Override
    public void onProviderEnabled(String s) {

    }

    @Override
    public void onProviderDisabled(String s) {

    }
}
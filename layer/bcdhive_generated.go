package layer

// Code generated by github.com/buildpacks/imgutil/tools/bcdhive_generator DO NOT EDIT

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io/ioutil"
)

const encodedBytes = "H4sIAAAAAAAC/+yaPWwbZRjH/5cmJAQoF/tcBanDoUSwcJXPvri+KaJN2gAlQZBWDBm4j9fkajuxbAOtqqCMGTuwMCB5YOhCxcYICxJiQB4Zu9EBoUgsgYEXPfeR88dZbRFDpT6/6HyP3/d5/T73PH/Feh+5LT6uKQpA14Oj/lcvff0zmZhBBNk6oheyd7CKVezgOjZxGTvooo0AHuqh3UQLO1jDOq7gTVzHNWyDeVr5Pqj+zVlgGIZhGIZhGIZ5Nth1g73QUNOxpA9Adk9KuVd/A/rrl75Mxn6YBaaS9QuAlFKSTffD+H4uY68XASwvL7/3/tb21oWNt27Q2E//SNmpA7SOrhcAnFcABdNfQAFUJVqrgfb8HfPhu0W8gikoURAzOoBXyZ5fpPeHY+MZ/ksT/Jdi/1W8POSvp/7q2Di0yD+KdSZHr/Na5Ds4NvFzh3zPhH2XEZ/hzwE2zrJ2GYZhGIZhGIZhmCc8/6vD5/9BovO/PnT+1wfmZQwdV++q0Rk+6QckvouxTWf7NdHx2kGrG+ynbYfGrgJdBfzyj78cSyl7KtCP4+lLKT+tq5imczpA9/B3CFeCdvMzpy3e3feDWiD8095F6D8b7nmsRr50vSNubzpNEc5fwmWsoTjyF3FvwvMm/Y4NLbvfcS0jd7MAttybwut28FBK2didOn1OdQH4Y/3+20cZ+9Fcsh/yj99feQ3AnYpT8cya6Rortm8Zpikcw7bKVcO2iiVRcVzftqsHad4PF4AP/lz9K6vONDdaZ5q4lRuv81xcm0fVeS4XPf8JPUsuqdd0WN8z0NWkvtu3W2KS/iiuuC2EWxPykyVninG9IZpir9tJxpK6JHH1coB9/7d7Wfv2ctn5OMqP5+MsgPMD+5qlImEl88fhvgoW88DSN899R/m4m0/1S2u/zaf6jeMO5118ggAN+GjBCX9708EFBNinWB6hpw8LT6gn2xdOtVLyjBXP9w1LXCwajueZRq1ccsuW5VsX7ZWDNI+Jnvp54Orn3atZeaS5rDz2tf+uq4+0VFe/asO6moL6WLrq51Nd9Qv/j65O49Im6+pEy87Hg0K2rjYH9i2VQ12VR3XVKgAn8zduUj4eFlJd0dq5c+O6ov97d1CBgwo8mKjBhAsDK7Dhw4IBEyYEHBiwYaGMamwVUYIIV7rwYcNGFQfp98fz/B3LMAzDMAzzNPFvAAAA///Odx8+ADAAAA=="

func BaseLayerBCD() ([]byte, error) {
	gzipBytes, err := base64.StdEncoding.DecodeString(encodedBytes)
	if err != nil {
		return nil, err
	}

	gzipReader, err := gzip.NewReader(bytes.NewBuffer(gzipBytes))
	if err != nil {
		return nil, err
	}

	decodedBytes, err := ioutil.ReadAll(gzipReader)
	if err != nil {
		return nil, err
	}

	return decodedBytes, nil
}

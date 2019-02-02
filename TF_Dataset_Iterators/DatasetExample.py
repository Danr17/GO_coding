# Demonstration of how to create Dataset with various supporting APIs
# Thanks to https://github.com/Prasad9/TF_Dataset_Iterators/

import tensorflow as tf
import numpy as np

# return an element at a time
dataset1 = tf.data.Dataset.from_tensor_slices(tf.range(10,15))

# return tuple ata a time
dataset2 = tf.data.Dataset.from_tensor_slices((tf.range(30,45,3), np.arange(60,70,2)))

# Code to try out data present in datasets
dataset = dataset2   # Change to required dataset
iterator = dataset.make_one_shot_iterator()
next_element = iterator.get_next()

with tf.Session() as sess:
  try:
    while True:
      print(sess.run(next_element))
  except:
    pass
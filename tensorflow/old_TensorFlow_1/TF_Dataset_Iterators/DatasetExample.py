# Demonstration of how to create Dataset with various supporting APIs
# Thanks to https://github.com/Prasad9/TF_Dataset_Iterators/

import tensorflow as tf
import numpy as np

# return single element at a time
dataset1 = tf.data.Dataset.from_tensor_slices(tf.range(10,15))

# return tuple at a time
dataset2 = tf.data.Dataset.from_tensor_slices((tf.range(30,45,3), np.arange(60,70,2)))

#making use of from_tensors
dataset3 = tf.data.Dataset.from_tensors(tf.range(10,15))

# return tuple of elements
dataset4 = tf.data.Dataset.from_tensors((tf.range(30,45,3), np.arange(60,70,2)))

dataset5 = tf.data.Dataset.from_tensors((tf.range(10), np.arange(5)))

#making use of generators
def generator(sequence_type):
  if sequence_type == 1:
    for i in range(5):
      yield 10 + i
  elif sequence_type == 2:
    for i in range(5):
      yield (30 + 3 * i, 60 + 2 * i)
  elif sequence_type == 3:
    for i in range(1,4):
      yield (i, ['Hi'] * i)

dataset6 = tf.data.Dataset.from_generator(generator, (tf.int32), args=([1]))
dataset7 = tf.data.Dataset.from_generator(generator, (tf.int32, tf.int32), args=([2]))
dataset8 = tf.data.Dataset.from_generator(generator, (tf.int32, tf.string), args=([3]))

# Code to try out data present in datasets
dataset = dataset7   # Change to required dataset
iterator = dataset.make_one_shot_iterator()
next_element = iterator.get_next()

with tf.Session() as sess:
  try:
    while True:
      print(sess.run(next_element))
  except:
    pass



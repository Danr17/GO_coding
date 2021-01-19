# Demonstration of Dataset Transformation
# Thanks to https://github.com/Prasad9/TF_Dataset_Iterators/

import tensorflow as tf

dataset = tf.data.Dataset.from_tensor_slices(tf.range(10))

# duplicate the dataset
dataset = dataset.repeat(2)

#Shuffle the dataset
dataset = dataset.shuffle(5)

#dataset map, multiply each element with 3
def map_fn(x):
    return x * 3
#dataset = dataset.map(map_fn)
dataset = dataset.map(lambda x: x * 3)

# filter out elements whose modulus 5 returns 1
def filter_fn(x):
    return tf.reshape(tf.not_equal(x % 5 ,1), [])
dataset = dataset.filter(filter_fn)

#batch at every 4 elements
dataset = dataset.batch(4)


# Code to try out data present in datasets
#dataset = dataset   # Change to required dataset
iterator = dataset.make_one_shot_iterator()
next_element = iterator.get_next()

with tf.Session() as sess:
  try:
    while True:
      print(sess.run(next_element))
  except:
    pass
from Infer import Infer

params = {
    'BATCH_SIZE': 64,
    'TOP_K': 5,

    'INFER_PATH': './Data/ImgsResize',
    'LABEL_PATH': './Data/classes.txt'
}
i = Infer(params)
i.infer()
# -*- mode: python ; coding: utf-8 -*-


a = Analysis(
    ['pkg/amqp_python/gateway.py'],
    pathex=['./pkg/amqp_python'],
    binaries=[],
    datas=[],
    hiddenimports=['proton', '_cffi_backend'],
    hookspath=[],
    hooksconfig={},
    runtime_hooks=[],
    excludes=[],
    noarchive=False,
)

# Add OpenSSL libraries if they're not automatically detected
openssl_libs = [('libssl.so.1.1', 'libssl.so.1.1', 'BINARY'), ('libcrypto.so.1.1', 'libcrypto.so.1.1', 'BINARY')]
a.binaries += openssl_libs

pyz = PYZ(a.pure)

exe = EXE(
    pyz,
    a.scripts,
    [],
    exclude_binaries=True,
    name='gateway',
    debug=True,
    bootloader_ignore_signals=False,
    strip=False,
    upx=True,
    console=True,
    disable_windowed_traceback=False,
    argv_emulation=False,
    target_arch=None,
    codesign_identity=None,
    entitlements_file=None,
)
coll = COLLECT(
    exe,
    a.binaries,
    a.datas,
    strip=False,
    upx=True,
    upx_exclude=[],
    name='gateway',
)
